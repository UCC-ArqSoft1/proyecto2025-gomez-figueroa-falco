import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import "./MisActividades.css";

const MisActividades = () => {
    const [misAct, setMisAct] = useState([]);     // siempre arranque como arreglo
    const [loading, setLoading] = useState(true);
    const userId = Number(localStorage.getItem("userId"));

    useEffect(() => {
        (async () => {
            try {
                const res = await fetch(`http://localhost:8080/misActividades/${userId}`);
                const data = await res.json();

                if (!res.ok) {
                    // Si el backend devolvió { error: "…" }
                    console.error("Error fetching misActividades:", data.error || data);
                    setMisAct([]);            // o manejar el mensaje de error como quieras
                } else if (Array.isArray(data)) {
                    setMisAct(data);          // aquí ya sabemos que es un arreglo
                } else {
                    console.warn("misActividades no es un array:", data);
                    setMisAct([]);            // caída suave
                }
            } catch (e) {
                console.error("Network error:", e);
                setMisAct([]);
            } finally {
                setLoading(false);
            }
        })();
    }, [userId]);

    if (loading) {
        return <div className="mis-actividades-page">Cargando mis actividades…</div>;
    }

    // Con la comprobación, podemos usar map sin miedo
    const lista = Array.isArray(misAct) ? misAct : [];

    return (
        <div className="mis-actividades-page">
            <h2>Mis actividades inscritas</h2>

            {lista.length === 0 ? (
                <p>No estás inscrito en ninguna actividad.</p>
            ) : (
                <div className="mis-actividades-list">
                    {lista.map(a => (
                        <div key={a.id} className="actividad-card">
                            <h3>{a.nombre}</h3>
                            {a.horarios?.length > 0 && (
                                <div className="actividad-horarios">
                                    {a.horarios.map(h => (
                                        <p key={h.id}>
                                            {h.dia}{" "}
                                            {new Date(h.hora_inicio).toLocaleTimeString([], {
                                                hour: "2-digit",
                                                minute: "2-digit",
                                            })}
                                            –{" "}
                                            {new Date(h.hora_fin).toLocaleTimeString([], {
                                                hour: "2-digit",
                                                minute: "2-digit",
                                            })}
                                        </p>
                                    ))}
                                </div>
                            )}
                            <small>Profesor: {a.profesor}</small>
                            <Link to={`/actividad/${a.id}`} className="detalle-btn">
                                Detalle
                            </Link>
                        </div>
                    ))}
                </div>
            )}

            <Link to="/actividades" className="back-btn">
                ← Volver a todas las actividades
            </Link>
        </div>
    );
};

export default MisActividades;
