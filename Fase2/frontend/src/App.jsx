import { useState } from "react";
import { Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import Estudiante from "./pages/Estudiante";
import Tutor from "./pages/Tutor";
import Administrador from "./pages/Administrador";
import "./App.css";
import TablaAlumnos from "./pages/TablaAlumnos";
import LibrosAdmin from "./pages/LibrosAdmin";
import A2 from "./pages/A2";
import AdminLibro from "./pages/AdminLibro";
import PublicacionTutor from "./pages/PublicacionTutor";
import EstudianteCursos from "./pages/EstudianteCursos";
import EstudianteLibros from "./pages/EstudianteLibros";
import EstudiantePublicacion from "./pages/EstudiantePublicacion";
import AdminReportes from "./pages/AdminReportes";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/principal/estudiante" element={<Estudiante />} />
        <Route path="/principal/estudiante/ver-cursos" element={<EstudianteCursos />} />
        <Route path="/principal/estudiante/ver-libros" element={<EstudianteLibros />} />
        <Route path="/principal/estudiante/ver-publicaciones" element={<EstudiantePublicacion />} />
        <Route path="/principal/tutor" element={<Tutor />} />
        <Route path="/principal/tutor/libro" element={<LibrosAdmin />} />
        <Route path="/principal/tutor/publicacion" element={<PublicacionTutor />} />
        <Route path="/principal/admin/cargar-archivos" element={<Administrador />} />
        <Route path="/principal/admin" element={<A2 />} />
        <Route path="/principal/admin/aceptar-libros" element={<AdminLibro />} />
        <Route path="/principal/admin/reportes" element={<AdminReportes />} />
        <Route path="/principal/admin/alumnos" element={<TablaAlumnos />} />

      </Routes>
    </>
  );
}

export default App;

