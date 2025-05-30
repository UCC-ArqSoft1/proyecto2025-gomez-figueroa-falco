
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import "./Login.css";



const Login = () => {

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [success, setSuccess] = useState("");
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();
        setError("");
        setSuccess("");
        try {
            const res = await fetch("http://localhost:8080/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ username, password })
            });
            const data = await res.json();

            if (res.ok && data.token) {
                localStorage.setItem("token", data.token);
                // decodificar y guardar userId y rol
                const payload = JSON.parse(atob(data.token.split('.')[1]));
                localStorage.setItem("userId", payload.userId);
                localStorage.setItem("rol", payload.rol); // Guarda el token si querés usarlo después
                setSuccess("¡Login correcto!");
                setError("");

                // Redirecciona, muestra mensaje o cambia de vista...
                navigate("/actividades");
            } else {
                setError(data.error || "Error desconocido");
                setSuccess("");
            }
        } catch (err) {
            setError("No se pudo conectar con el servidor");
            setSuccess("");
        }
    };
    return (
        <div className="login-container">
            <form className="login-form" onSubmit={handleLogin}>
                <h2>
                    Iniciar Sesion
                </h2>
                <input
                    type="text"
                    placeholder="Usuario"
                    onChange={(e) => setUsername(e.target.value)}
                    value={username}
                    required
                />
                <input
                    type="password"
                    placeholder="Contraseña"
                    onChange={(e) => setPassword(e.target.value)}
                    value={password}
                    required
                />
                <button type="submit">Ingresar</button>
                {error && <div className="login-error">{error}</div>}
                {success && <div className="login-success">{success}</div>}
            </form>
        </div>
    )
}
export default Login;
