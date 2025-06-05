import { useState, useEffect, useCallback } from "react";
import { Link } from "react-router-dom";
import "./Actividades.css";

/** debounce simple: espera N ms antes de lanzar la búsqueda */
function useDebounce(value, delay = 400) {
    const [debounced, setDebounced] = useState(value);
    useEffect(() => {
        const id = setTimeout(() => setDebounced(value), delay);
        return () => clearTimeout(id);
    }, [value, delay]);
    return debounced;
}

export default function Actividades() {
    const [actividades, setActividades] = useState([]);
    const [loading, setLoading] = useState(false);
    const [busqueda, setBusqueda] = useState("");

    const rol = localStorage.getItem("rol") || "";

    /* 1️⃣  400 ms después de que el usuario deja de teclear
           hacemos la petición con ?q=… */
    const termino = useDebounce(busqueda);

    const fetchActividades = useCallback(async () => {
        setLoading(true);
        try {
            const url = termino.trim()
                ? `http://localhost:8080/actividades?q=${encodeURIComponent(termino)}`
                : "http://localhost:8080/actividades";
            const res = await fetch(url);
            const data = await res.json();
            setActividades(data);
        } catch (err) {
            console.error("Error obteniendo actividades:", err);
            setActividades([]);
        } finally {
            setLoading(false);
        }
    }, [termino]);

    /* 2️⃣  dispara al montar el componente y cada vez que cambia ‘termino’ */
    useEffect(() => {
        fetchActividades();
    }, [fetchActividades]);

    /* ---------- UI ---------- */
    return (
        <div className="actividades-page">
            <header className="busqueda-header">
                <input
                    className="busqueda-input"
                    type="text"
                    placeholder="Buscar por título, horario o categoría…"
                    value={busqueda}
                    onChange={(e) => setBusqueda(e.target.value)}
                />

                {rol === "SOCIO" && (
                    <Link to="/mis-actividades" className="mis-actividades-link">
                        Mis actividades
                    </Link>
                )}
            </header>

            {loading && <p className="cargando">Cargando…</p>}

            <section className="actividades-container">
                {actividades.map((a) => (
                    <div key={a.id} className="actividad-card">
                        <h3>{a.nombre}</h3>

                        {a.horarios?.length > 0 && (
                            <div className="actividad-horarios">
                                {a.horarios.map((h) => (
                                    <p key={h.id}>
                                        {h.dia}{" "}
                                        {new Date(h.hora_inicio).toLocaleTimeString([], {
                                            hour: "2-digit",
                                            minute: "2-digit",
                                        })}{" "}
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

                {!loading && actividades.length === 0 && (
                    <p className="sin-resultados">No se encontraron actividades.</p>
                )}
            </section>
        </div>
    );
}
