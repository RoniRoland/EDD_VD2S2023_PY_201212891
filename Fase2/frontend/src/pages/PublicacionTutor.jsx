import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function PublicacionTutor() {

    const [contenidoPublicacion, setContenidoPublicacion] = useState("");
    const CargarPublicacionTutor = async (e) => {
        e.preventDefault();
        const valorLocal = localStorage.getItem("user");
        const response = await fetch(
            "http://localhost:4000/registrar-publicacion",
            {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    Carnet: parseInt(valorLocal),
                    Nombre: valorLocal,
                    Contenido: contenidoPublicacion,
                }),
            }
        );

        const result = await response.json();
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
                    <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Carga de Publicaciones</h1>
                    <br />
                    <br />
                    <div class="form-floating">
                        <textarea class="form-control" placeholder="Leave a comment here" id="floatingTextarea" value={contenidoPublicacion}
                            onChange={(e) => setContenidoPublicacion(e.target.value)}></textarea>
                        <label for="floatingTextarea">Area de texto</label>
                    </div>
                    <br />
                    <br />
                    <button type="button" class="btn btn-success" onClick={CargarPublicacionTutor}>Cargar Publicacion</button>
                    <br />
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-primary" onClick={salir}>
                            Salir
                        </button>
                    </center>
                </form>
            </div>
        </div>
    )
}

export default PublicacionTutor