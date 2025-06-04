import { useParams, Link } from "react-router-dom";
import { useState, useEffect } from "react";
import "./ActividadDetalle.css";

const ActividadDetalle = () => {
    const { id } = useParams();
    const [act, setAct] = useState(null);
    const [loading, setLoad] = useState(true);

    const userId = Number(localStorage.getItem("userId"));
    const rol = localStorage.getItem("rol");

    useEffect(() => {
        fetch(`http://localhost:8080/actividad/${id}`)
            .then(r => r.json())
            .then(data => {
                setAct(data);
                setLoad(false);
            })
            .catch(() => setLoad(false));
    }, [id]);

    const handleInscribirse = async (horarioId, dia) => {
        try {
            const res = await fetch(
                `http://localhost:8080/inscripcion`,
                {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({
                        userId,
                        actividadId: act.id,
                        horarioId,
                        dia
                    })
                }
            );
            const data = await res.json();
            if (res.ok) {
                alert(data.msg);
            } else {
                alert(data.error);
            }
        } catch (e) {
            console.error(e);
            alert("Error al inscribirse");
        }
    };

    if (loading) return <div>Cargando detalle…</div>;
    if (!act) return <div>Actividad no encontrada</div>;

    console.log("ROL desde localStorage:", rol);

    return (
        <div className="actividad-detalle-page">
            <div className="actividad-detalle-card">
                <h2 className="detalle-title">{act.nombre}</h2>
                <img className="detalle-img" src={act.imagen} alt={act.nombre} />
                <p className="detalle-text">{act.descripcion}</p>
                <p><strong>Categoría:</strong> {act.categoria}</p>
                <p><strong>Profesor:</strong> {act.profesor}</p>
                <p><strong>Cupo:</strong> {act.cupo_total}</p>

                {act.horarios?.length > 0 && (
                    <div className="detalle-horarios">
                        <h3>Horarios</h3>
                        {act.horarios.map(h => (
                            <div key={h.id} className="detalle-horario-item">
                                <p>
                                    {h.dia}{" "}
                                    {new Date(h.hora_inicio).toLocaleTimeString([], {
                                        hour: "2-digit",
                                        minute: "2-digit",
                                    })} –{" "}
                                    {new Date(h.hora_fin).toLocaleTimeString([], {
                                        hour: "2-digit",
                                        minute: "2-digit",
                                    })}
                                </p>
                                {rol === "SOCIO" && (
                                    <button
                                        className="inscribir-btn"
                                        onClick={() => handleInscribirse(h.id, h.dia)}
                                    >
                                        Inscribirse
                                    </button>
                                )}
                            </div>
                        ))}
                    </div>
                )}

                <Link to="/actividades" className="detalle-back">
                    ← Volver
                </Link>
            </div>
        </div>
    );
};

export default ActividadDetalle;
