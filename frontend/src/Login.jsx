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

        if (!username || !password) {
            setError("Por favor, completa todos los campos");
            return;
        }

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
                const payload = JSON.parse(atob(data.token.split('.')[1]));
                localStorage.setItem("userId", payload.userId);
                localStorage.setItem("rol", payload.rol);
                setSuccess("¡Login correcto!");
                setError("");
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
            <form className="login-form" onSubmit={handleLogin} noValidate>
                <h2>Iniciar Sesión</h2>
                <input
                    type="text"
                    placeholder="Usuario"
                    onChange={(e) => setUsername(e.target.value)}
                    value={username}
                    required
                    aria-label="Usuario"
                    aria-required="true"
                />
                <input
                    type="password"
                    placeholder="Contraseña"
                    onChange={(e) => setPassword(e.target.value)}
                    value={password}
                    required
                    aria-label="Contraseña"
                    aria-required="true"
                />
                <button type="submit">Ingresar</button>
                {error && <div className="login-error">{error}</div>}
                {success && <div className="login-success">{success}</div>}
            </form>
        </div>
    );
};

export default Login;
