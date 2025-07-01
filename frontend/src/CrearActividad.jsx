// src/pages/CrearActividad.jsx
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./CrearActividad.css";

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

export default function CrearActividad() {
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
    const navigate = useNavigate();

    const handleChange = (e) => {
        const { name, value } = e.target;
        setForm(f => ({ ...f, [name]: value }));
        setError("");
        if (name === "cupo_total") {
            setHorarios(hs => hs.map(h => ({ ...h, cupo_horario: value })));
        }
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
        setHorarios(hs => [
            ...hs,
            { hora_inicio: "", hora_fin: "", cupo_horario: form.cupo_total }
        ]);
    };
    const removeHorario = (idx) => {
        setHorarios(hs => hs.filter((_, i) => i !== idx));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        const data = new FormData();
        Object.entries(form).forEach(([k, v]) => v && data.append(k, v));
        // Transformar horarios al formato esperado por el backend
        const horariosTransformados = horarios.map(h => ({
            dia: getDiaSemana(h.hora_inicio),
            hora_inicio: toFechaHora(h.hora_inicio),
            hora_fin: toFechaHora(h.hora_fin),
            cupo_horario: Number(form.cupo_total)
        }));
        data.append("horarios", JSON.stringify(horariosTransformados));
        if (file) data.append("imagen", file);

        try {
            const token = localStorage.getItem("token");
            const res = await fetch("http://localhost:8080/actividades", {
                method: "POST",
                headers: {
                    Authorization: `Bearer ${token}`,
                },
                body: data,
            });
            const result = await res.json();
            if (res.ok) {
                navigate("/actividades");
            } else {
                setError(result.error || "Error creando actividad");
            }
        } catch (e) {
            console.error(e);
            setError("Fallo de red");
        }
    };

    return (
        <div className="crear-actividad-page">
            <h2>Crear Nueva Actividad</h2>
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
                <input
                    type="file"
                    name="imagen"
                    accept="image/*"
                    onChange={handleFile}
                />

                {error && <div className="form-error">{error}</div>}
                <div className="form-actions">
                    <button type="button" onClick={() => navigate('/actividades')} className="cancel-btn">
                        ← Volver
                    </button>
                    <button type="submit">Crear Actividad</button>
                </div>
            </form>
        </div>
    );
}
