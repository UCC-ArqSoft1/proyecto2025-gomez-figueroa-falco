import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Login from "./Login";
import Actividades from "./Actividades"; // lo vas a crear ahora

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/actividades" element={<Actividades />} />
      </Routes>
    </Router>
  );
}

export default App;
