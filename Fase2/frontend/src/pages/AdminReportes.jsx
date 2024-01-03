import React, { useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function AdminReportes() {
    const [imagen, setImagen] = useState(
        "https://i.imgur.com/STNQphI.png"
    );

    const salir = (e) => {
        e.preventDefault();
        console.log("Listo");
        window.open("/principal/admin", "_self");
    };

    const validar = (data) => {
        console.log(data);
        setImagen(data.imagen.Imagenbase64);
    };

    const reporteGrafo = async (e) => {
        e.preventDefault();
        fetch("http://localhost:4000/reporte-grafo", {})
            .then((response) => response.json())
            .then((data) => validar(data));
    };

    const reporteArbol = async (e) => {
        e.preventDefault();
        fetch("http://localhost:4000/reporte-arbol", {})
            .then((response) => response.json())
            .then((data) => validar(data));
    };

    const reporteBlockchain = async (e) => {
        e.preventDefault();
        fetch("http://localhost:4000/reporte-merkle", {})
            .then((response) => response.json())
            .then((data) => validar(data));
    };

    return (
        <div className="form-signin1">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Menu de Reportes</h1>
                    <br />
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-primary" onClick={reporteArbol}>
                            Reporte Arbol B
                        </button>
                    </center>
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-primary" onClick={reporteGrafo}>
                            Reporte Grafo
                        </button>
                    </center>
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-primary" onClick={reporteBlockchain}>
                            Reporte Arbol Merkle
                        </button>
                    </center>
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-success" onClick={salir}>
                            Salir
                        </button>
                    </center>
                    <br />
                    <br />
                    <center>
                        <img src={imagen} width="500" height="500" alt="some value" />
                    </center>

                </form>
            </div>
        </div>
    )
}

export default AdminReportes