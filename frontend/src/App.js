import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./Home";              // ← Nuevo
import Login from "./Login";
import Actividades from "./Actividades";
import ActividadDetalle from "./ActividadDetalle";
import MisActividades from "./MisActividades";
import "./App.css";

function App() {
  return (
    <Router>
      <Routes>
        {/* Pantalla inicial (Home) */}
        <Route path="/" element={<Home />} />

        {/* Página de Login */}
        <Route path="/login" element={<Login />} />

        {/* El resto de rutas de la app */}
        <Route path="/actividades" element={<Actividades />} />
        <Route path="/actividad/:id" element={<ActividadDetalle />} />
        <Route path="/mis-actividades" element={<MisActividades />} />
      </Routes>
    </Router>
  );
}

export default App;
