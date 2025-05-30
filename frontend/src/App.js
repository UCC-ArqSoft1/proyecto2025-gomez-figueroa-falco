import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Actividades from "./Actividades";
import './App.css';
import ActividadDetalle from "./ActividadDetalle";
import MisActividades from "./MisActividades";


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/actividades" element={<Actividades />} />
        <Route path="/actividad/:id" element={<ActividadDetalle />} />
        <Route path="/mis-actividades" element={<MisActividades />} />
      </Routes>
    </Router>
  );
}

export default App;