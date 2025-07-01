import { useParams, Link } from "react-router-dom";
import { useState, useEffect } from "react";
import "./ActividadDetalle.css";

const CustomAlert = ({ message, onClose }) => (
    <div className="alert-overlay">
        <div className="alert-container">
            <p className="alert-message">{message}</p>
            <button className="alert-button" onClick={onClose}>
                Aceptar
            </button>
        </div>
    </div>
);

const ActividadDetalle = () => {
    const { id } = useParams();
    const [act, setAct] = useState(null);
    const [loading, setLoad] = useState(true);
    const [alertMessage, setAlertMessage] = useState("");
    const [showDeleteConfirm, setShowDeleteConfirm] = useState(false);
    const [deleteSuccess, setDeleteSuccess] = useState(false);

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
                setAlertMessage(data.msg);
                setAct(prev => ({
                    ...prev,
                    horarios: prev.horarios.map(h =>
                        h.id === horarioId
                            ? { ...h, cupo_horario: h.cupo_horario - 1 }
                            : h
                    )
                }));
            } else {
                setAlertMessage(data.error);
            }
        } catch (e) {
            console.error(e);
            setAlertMessage("Error al inscribirse");
        }
    };

    const handleEliminarActividad = () => {
        setShowDeleteConfirm(true);
    };

    const confirmEliminar = async () => {
        try {
            const token = localStorage.getItem("token");
            const res = await fetch(`http://localhost:8080/actividades/${id}`, {
                method: "DELETE",
                headers: { Authorization: `Bearer ${token}` },
            });
            if (res.ok) {
                setDeleteSuccess(true);
                setTimeout(() => {
                    window.location.href = "/actividades";
                }, 1500);
            } else {
                const data = await res.json();
                setAlertMessage(data.error || "Error eliminando actividad");
            }
        } catch (e) {
            setAlertMessage("Error de red");
        }
        setShowDeleteConfirm(false);
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
                                <div className="detalle-horario-info">
                                    <span>
                                        {h.dia} {new Date(h.hora_inicio).toLocaleTimeString([], {
                                            hour: "2-digit",
                                            minute: "2-digit",
                                        })} – {new Date(h.hora_fin).toLocaleTimeString([], {
                                            hour: "2-digit",
                                            minute: "2-digit",
                                        })}
                                    </span>
                                    <span className="detalle-cupo-horario" style={{ marginLeft: 16, color: '#888', fontSize: '1rem' }}>
                                        <strong>Cupo disponible:</strong> {h.cupo_horario}
                                    </span>
                                </div>
                                {(rol === "SOCIO" || rol === "ADMIN") && (
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

                <div className="detalle-actions">
                    <Link to="/actividades" className="detalle-back">
                        ← Volver
                    </Link>
                    {rol === "ADMIN" && (
                        <div className="detalle-actions-right">
                            <Link to={`/editar-actividad/${id}`} className="editar-btn">
                                ✏️ Editar Actividad
                            </Link>
                            <button className="eliminar-btn" onClick={handleEliminarActividad}>
                                🗑️ Eliminar
                            </button>
                        </div>
                    )}
                </div>
            </div>
            {alertMessage && (
                <CustomAlert
                    message={alertMessage}
                    onClose={() => setAlertMessage("")}
                />
            )}
            {showDeleteConfirm && (
                <div className="alert-overlay">
                    <div className="alert-container">
                        <p className="alert-message">¿Estás seguro de eliminar la actividad? Esta acción no se puede deshacer.</p>
                        <button className="alert-button" onClick={confirmEliminar} style={{ background: '#c1121f', color: 'white' }}>Sí, eliminar</button>
                        <button className="alert-button" onClick={() => setShowDeleteConfirm(false)}>Volver</button>
                    </div>
                </div>
            )}
            {deleteSuccess && (
                <CustomAlert message="Actividad borrada" onClose={() => window.location.href = "/actividades"} />
            )}
        </div>
    );
};

export default ActividadDetalle;
