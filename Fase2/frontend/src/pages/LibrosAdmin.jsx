import React, { useState } from "react";

function LibrosAdmin() {
  const [contenidoPDF, setContenidoPDF] = useState("");

  const uploadFileTutor = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();

    reader.onload = async (event) => {
      const content = event.target.result;
      console.log(content);
      setContenidoPDF(content);
      const valorLocal = localStorage.getItem("user");
      const response = await fetch("http://localhost:4000/registrar-libro", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          Carnet: parseInt(valorLocal),
          Nombre: file.name,
          Contenido: content,
        }),
      });

      const result = await response.json();
    };

    reader.onerror = (error) => {
      console.error("Error al leer el archivo:", error);
    };

    reader.readAsDataURL(file);
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/principal/tutor", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Carga de Libros</h1>
          <br />
          <br />
          <div className="input-group mb-3">
            <label className="input-group-text">Cargar Tutores</label>
            <input
              className="form-control"
              id="inputGroupFile01"
              type="file"
              accept=".pdf"
              onChange={uploadFileTutor}
            />
          </div>
          <div className="mb-3 fw-normal">
            <iframe src={contenidoPDF} width="800" height="800" />
          </div>
          <center>
            <button className="w-50 btn btn-outline-primary" onClick={salir}>
              Salir
            </button>
          </center>
        </form>
      </div>
    </div>
  );
}

export default LibrosAdmin;
