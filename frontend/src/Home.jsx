// src/pages/Home.jsx
import { Link } from "react-router-dom";
import "./Home.css";

export default function Home() {
    return (
        <div className="home-page">
            <div className="home-content">
                <h1 className="home-title">Bienvenido a Nuestro Gimnasio</h1>
                <p className="home-subtitle">¡Ponte en forma con nosotros!</p>
                <Link to="/login" className="home-button">
                    Ir al Login
                </Link>
            </div>
        </div>
    );
}
