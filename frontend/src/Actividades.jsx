import { useState, useEffect } from "react";
import { Link } from "react-router-dom"
import './Actividades.css';

const Actividades = () => {
    const [actividades, setActividades] = useState([]);
    const [loading, setLoading] = useState(true);
    const [busqueda, setBusqueda] = useState("");

    useEffect(() => {
        fetch("http://localhost:8080/actividades")
            .then(res => res.json())
            .then(data => {
                setActividades(data);
                setLoading(false);
            })
            .catch(() => setLoading(false));
    }, []);

    if (loading) return <div>Cargando actividades...</div>;


    return (
        <div className="actividades-page">
            <div className="busqueda-container">
                <input
                    type="text"
                    className="busqueda-input"
                    placeholder="Buscar actividad por título, horario o profesor"
                    value={busqueda}
                    onChange={e => setBusqueda(e.target.value)}
                />
            </div>

            <div className="actividades-container">
                {actividades.map(a => (
                    <div key={a.id} className="actividad-card">
                        <h3>{a.nombre}</h3>

                        {/* mostramos todos los horarios */}
                        {a.horarios?.length > 0 && (
                            <div className="actividad-horarios">
                                {a.horarios.map(h => (
                                    <p key={h.id}>
                                        {h.dia}{" "}
                                        {new Date(h.hora_inicio).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                                        –
                                        {new Date(h.hora_fin).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                                    </p>
                                ))}
                            </div>
                        )}

                        <small>Profesor: {a.profesor}</small>
                        {/* botón detalle */}
                        <Link to={`/actividad/${a.id}`} className="detalle-btn">
                            Detalle           </Link>
                    </div>
                ))}

                {actividades.length === 0 && (
                    <p className="sin-resultados">No hay actividades que coincidan.</p>
                )}
            </div>
        </div>
    );
};

export default Actividades;
