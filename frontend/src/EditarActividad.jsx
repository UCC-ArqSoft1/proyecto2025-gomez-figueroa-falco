import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import "./EditarActividad.css";

const diasSemana = [
    "Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"
];

function getDiaSemana(fechaStr) {
    const date = new Date(fechaStr);
    return diasSemana[date.getDay()];
}

function toFechaHora(fechaStr) {
    // Convierte '2025-07-05T18:51' a '2025-07-05 18:51'
    return fechaStr.replace("T", " ");
}

function toDateTimeLocal(dateStr) {
    // Soporta formatos con o sin segundos y zona horaria
    if (!dateStr) return "";
    // Si ya está en formato YYYY-MM-DDTHH:MM, devolver igual
    if (/^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}$/.test(dateStr)) return dateStr;
    // Si viene con espacio, reemplazar por T
    let base = dateStr.replace(" ", "T");
    // Si tiene segundos o zona horaria, recortar
    // Ejemplo: 2025-07-05T18:51:00Z o 2025-07-05T18:51:00
    const match = base.match(/^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2})/);
    return match ? match[1] : base;
}

export default function EditarActividad() {
    const { id } = useParams();
    const navigate = useNavigate();

    const [form, setForm] = useState({
        nombre: "",
        descripcion: "",
        categoria: "",
        profesor: "",
        cupo_total: "",
    });
    const [horarios, setHorarios] = useState([
        { hora_inicio: "", hora_fin: "", cupo_horario: "" }
    ]);
    const [file, setFile] = useState(null);
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(true);

    // Cargar datos de la actividad
    useEffect(() => {
        const cargarActividad = async () => {
            try {
                const res = await fetch(`http://localhost:8080/actividad/${id}`);
                if (!res.ok) {
                    setError("Actividad no encontrada");
                    setLoading(false);
                    return;
                }

                const actividad = await res.json();

                // Llenar el formulario con los datos existentes
                setForm({
                    nombre: actividad.nombre || "",
                    descripcion: actividad.descripcion || "",
                    categoria: actividad.categoria || "",
                    profesor: actividad.profesor || "",
                    cupo_total: actividad.cupo_total?.toString() || "",
                });

                // Convertir horarios al formato del formulario
                if (actividad.horarios && actividad.horarios.length > 0) {
                    const horariosFormateados = actividad.horarios.map(h => ({
                        hora_inicio: toDateTimeLocal(h.HoraInicio || h.hora_inicio),
                        hora_fin: toDateTimeLocal(h.HoraFin || h.hora_fin),
                        cupo_horario: (h.CupoHorario || h.cupo_horario)?.toString() || ""
                    }));
                    setHorarios(horariosFormateados);
                }

                setLoading(false);
            } catch (err) {
                console.error("Error cargando actividad:", err);
                setError("Error al cargar la actividad");
                setLoading(false);
            }
        };

        cargarActividad();
    }, [id]);

    const handleChange = (e) => {
        setForm(f => ({ ...f, [e.target.name]: e.target.value }));
        setError("");
    };

    const handleFile = (e) => {
        setFile(e.target.files[0]);
        setError("");
    };

    const handleHorarioChange = (idx, e) => {
        const { name, value } = e.target;
        setHorarios(hs => hs.map((h, i) => i === idx ? { ...h, [name]: value } : h));
    };

    const addHorario = () => {
        setHorarios(hs => [...hs, { hora_inicio: "", hora_fin: "", cupo_horario: "" }]);
    };

    const removeHorario = (idx) => {
        setHorarios(hs => hs.filter((_, i) => i !== idx));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        // Transformar horarios al formato esperado por el backend
        const horariosTransformados = horarios.map(h => ({
            dia: getDiaSemana(h.hora_inicio),
            hora_inicio: toFechaHora(h.hora_inicio),
            hora_fin: toFechaHora(h.hora_fin),
            cupo_horario: Number(h.cupo_horario)
        }));

        const datosActividad = {
            nombre: form.nombre,
            descripcion: form.descripcion,
            categoria: form.categoria,
            profesor: form.profesor,
            cupo_total: Number(form.cupo_total),
            imagen: "", // Por ahora no manejamos cambio de imagen
            horarios: horariosTransformados
        };

        try {
            const token = localStorage.getItem("token");
            const res = await fetch(`http://localhost:8080/actividades/${id}`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${token}`,
                },
                body: JSON.stringify(datosActividad),
            });

            const result = await res.json();
            if (res.ok) {
                navigate(`/actividad/${id}`);
            } else {
                setError(result.error || "Error editando actividad");
            }
        } catch (e) {
            console.error(e);
            setError("Fallo de red");
        }
    };

    if (loading) {
        return <div className="crear-actividad-page">Cargando actividad...</div>;
    }

    return (
        <div className="crear-actividad-page">
            <h2>Editar Actividad</h2>
            <form className="crear-actividad-form" onSubmit={handleSubmit}>
                <input
                    name="nombre"
                    placeholder="Nombre"
                    value={form.nombre}
                    onChange={handleChange}
                    required
                />
                <textarea
                    name="descripcion"
                    placeholder="Descripción"
                    value={form.descripcion}
                    onChange={handleChange}
                    required
                />
                <input
                    name="categoria"
                    placeholder="Categoría"
                    value={form.categoria}
                    onChange={handleChange}
                    required
                />
                <input
                    name="profesor"
                    placeholder="Profesor"
                    value={form.profesor}
                    onChange={handleChange}
                    required
                />
                <input
                    name="cupo_total"
                    type="number"
                    placeholder="Cupo Total"
                    value={form.cupo_total}
                    onChange={handleChange}
                    required
                />

                <hr />
                <h3>Horarios</h3>
                {horarios.map((h, idx) => (
                    <div key={idx} className="horario-block">
                        <input
                            name="hora_inicio"
                            type="datetime-local"
                            placeholder="Fecha y hora de inicio"
                            value={h.hora_inicio}
                            onChange={e => handleHorarioChange(idx, e)}
                            required
                        />
                        <input
                            name="hora_fin"
                            type="datetime-local"
                            placeholder="Fecha y hora de fin"
                            value={h.hora_fin}
                            onChange={e => handleHorarioChange(idx, e)}
                            required
                        />
                        <input
                            name="cupo_horario"
                            type="number"
                            placeholder="Cupo Horario"
                            value={h.cupo_horario}
                            onChange={e => handleHorarioChange(idx, e)}
                            required
                        />
                        {horarios.length > 1 && (
                            <button type="button" onClick={() => removeHorario(idx)}>
                                Quitar
                            </button>
                        )}
                    </div>
                ))}
                <button type="button" onClick={addHorario} className="add-horario-btn">
                    + Agregar Horario
                </button>

                <hr />
                <h3>Imagen</h3>
                <p className="editar-img-aviso">
                    La funcionalidad de cambio de imagen no está disponible en esta versión.
                </p>

                {error && <div className="form-error">{error}</div>}

                <div className="form-actions">
                    <button type="button" onClick={() => navigate(`/actividad/${id}`)} className="cancel-btn">
                        ← Volver
                    </button>
                    <button type="submit">Guardar Cambios</button>
                </div>
            </form>
        </div>
    );
} 