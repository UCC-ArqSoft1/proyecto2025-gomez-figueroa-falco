import { useState, useEffect } from "react";
import './Actividades.css';

const Actividades = () => {
    const [actividades, setActividades] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetch("http://localhost:8080/actividades")
            .then((res) => res.json())
            .then((data) => {
                setActividades(data);
                setLoading(false);
            })
            .catch((err) => {
                setLoading(false);
            });
    }, []);

    if (loading) return <div>Cargando actividades...</div>;

    return (
        <div className="actividades-page">
            <div className="actividades-container">
                {actividades.map(a => (
                    <div key={a.id} className="actividad-card">
                        <h3>{a.nombre}</h3>
                        <p>{a.descripcion}</p>
                        <small>
                            Categoría: {a.categoria} | Profesor: {a.profesor}
                        </small>
                    </div>
                ))}
            </div>
        </div>
    );

};

export default Actividades;