import { useState, useEffect } from "react";
import "./MisActividades.css";

const MisActividades = () => {
    const [misAct, setMisAct] = useState([]);
    const [loading, setLoading] = useState(true);
    const userId = Number(localStorage.getItem("userId"));

    useEffect(() => {
        (async () => {
            try {
                const res = await fetch(`http://localhost:8080/misActividades/${userId}`);
                const data = await res.json();
                setMisAct(Array.isArray(data) ? data : []);
            } catch (e) {
                console.error(e);
                setMisAct([]);
            } finally {
                setLoading(false);
            }
        })();
    }, [userId]);

    if (loading) {
        return <div className="mis-actividades-page">Cargando mis actividades…</div>;
    }

    return (
        <div className="mis-actividades-page">
            <h2>Actividades a las que estoy inscripcto</h2>
            {misAct.length === 0 ? (
                <p>No estás inscripto en ninguna actividad.</p>
            ) : (
                <div className="mis-actividades-list">
                    {misAct.map(a => (
                        <div key={a.Id} className="actividad-card">
                            {/* Usamos PascalCase que proviene del JSON GORM */}
                            <h3 className="actividad-nombre">{a.Nombre}</h3>
                            <p className="actividad-profesor">
                                <strong>Profesor:</strong> {a.Profesor}
                            </p>

                            {/* Horarios también en PascalCase */}
                            {a.Horarios?.length > 0 && (
                                <div className="actividad-horarios">
                                    {a.Horarios.map(h => (
                                        <p key={h.Id} className="horario-item">
                                            <strong>{h.Dia}:</strong>{" "}
                                            {new Date(h.HoraInicio).toLocaleTimeString([], {
                                                hour: "2-digit",
                                                minute: "2-digit",
                                            })} –{" "}
                                            {new Date(h.HoraFin).toLocaleTimeString([], {
                                                hour: "2-digit",
                                                minute: "2-digit",
                                            })}
                                        </p>
                                    ))}
                                </div>
                            )}
                        </div>
                    ))}
                </div>
            )}
            <button
                className="back-btn"
                onClick={() => (window.location.href = "/actividades")}
            >
                ← Volver a todas las actividades
            </button>
        </div>
    );
};

export default MisActividades;
