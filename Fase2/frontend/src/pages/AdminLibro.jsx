import React, { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import '../styles/otros.css'

function AdminLibro() {

    const [libros, setLibros] = useState([]);
    const [eleccion, setEleccion] = useState(0);
    const [imagen, setImagen] = useState(
        "https://i.imgur.com/qf0Sh0j.jpg"
    );

    useEffect(() => {
        async function PedirLibros() {
            const response = await fetch("http://localhost:4000/enviar-libros-admin");
            const result = await response.json();
            console.log(result);
            if (result.status === 200) {
                setLibros(result.Arreglo);
            }
        }
        PedirLibros();
    }, []);

    const handleChange = (e) => {
        var j = parseInt(e.target.value);
        setEleccion(j);
        console.log(j);
    };

    const LibrosObtenidos = () => {
        console.log(libros);
        return (
            <div>
                <select
                    className="form-control"
                    aria-label=".form-select-lg example"
                    onChange={handleChange}
                >
                    <option value={0}>Elegir Libro....</option>
                    {libros.map((item, j) => (
                        <option value={j} key={j}>
                            {item.Nombre}
                        </option>
                    ))}
                </select>
            </div>
        );
    };

    const LibrosDefault = () => {
        console.log(libros);
        return (
            <div>
                <select
                    className="form-control"
                    aria-label=".form-select-lg example"
                    onChange={handleChange}
                >
                    <option value={0}>Elegir Libro....</option>
                </select>
            </div>
        );
    };

    const aceptar = async (e) => {
        e.preventDefault();
        const valorLocal = localStorage.getItem("user");
        if (libros.length > 0) {
            const response = await fetch("http://localhost:4000/registrar-log", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    Accion: "Aceptado",
                    Nombre: libros[eleccion].Nombre,
                    Tutor: libros[eleccion].Tutor,
                    Curso: libros[eleccion].Curso,
                }),
            });

            const result = await response.json();
        }
    };

    const rechazar = async (e) => {
        e.preventDefault();
        const valorLocal = localStorage.getItem("user");
        if (libros.length > 0) {
            const response = await fetch("http://localhost:4000/registrar-log", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    Accion: "Rechazado",
                    Nombre: libros[eleccion].Nombre,
                    Tutor: libros[eleccion].Tutor,
                }),
            });

            const result = await response.json();
        }
    };

    const finalizar = async (e) => {
        e.preventDefault();
        const response = await fetch("http://localhost:4000/finalizar-libros");
        const result = await response.json();
    };





    const salir = (e) => {
        e.preventDefault();
        window.open("/principal/admin", "_self");
    };



    return (

        <div className="form-signin1">
            <div className="text-center">
                <form className="card card-body">
                    <h1 className="display-4 fw-bold text-body-emphasis Menu-carga">Aceptar Libros {localStorage.getItem("empleado")}</h1>
                    <br />
                    <br />
                    <h2>Libros</h2>
                    <br />
                    <br />
                    <div className="col align-self-center">
                        {libros.length > 0 ? <LibrosObtenidos /> : <LibrosDefault />}
                    </div>
                    <br />
                    <br />
                    <center>
                        <img src={imagen} width="400" height="500" alt="some value" />
                    </center>
                    <br />
                    <br />


                    <br />
                    <div className="row align-items-start">
                        <div className="col">
                            <center>
                                <button
                                    className="w-50 btn btn-danger btn-lg"
                                    onClick={rechazar}
                                >
                                    Rechazar
                                </button>
                            </center>
                        </div>
                        <div className="col">
                            <center>
                                <button
                                    className="w-50 btn btn-success btn-lg"
                                    onClick={aceptar}
                                >
                                    Aceptar
                                </button>
                            </center>
                        </div>
                        <div className="col">
                            <center>
                                <button className="w-50 btn btn-warning btn-lg" onClick={finalizar}>
                                    Finalizar
                                </button>
                            </center>
                        </div>
                    </div>
                    <br />
                    <center>
                        <button className="w-50 btn btn-outline-primary btn-lg" onClick={salir}>
                            Salir
                        </button>
                    </center>


                </form>
            </div>
        </div>
    )
}

export default AdminLibro