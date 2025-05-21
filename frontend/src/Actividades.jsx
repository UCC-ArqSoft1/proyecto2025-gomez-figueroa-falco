import { useState, useEffect } from "react";

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
        <div>
            <h2>Lista de Actividades</h2>
            <ul>
                {actividades.map((a) => (
                    <li key={a.id}>
                        <strong>{a.nombre}</strong> - {a.descripcion} <br />
                        Categoría: {a.categoria} | Profesor: {a.profesor}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Actividades;
