import React, { useState } from 'react';
import '../styles/login.css'
import "bootstrap/dist/css/bootstrap.min.css";

function Login() {
    const [isChecked, setIsChecked] = useState(false);
    const [userName, setUserName] = useState("");
    const [passwordUser, setPasswordUser] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch("http://localhost:4000/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                UserName: userName,
                Password: passwordUser,
                Tutor: isChecked,
            }),
        });

        const result = await response.json();
        if (result.rol == 0) {
            alert("Credenciales Incorrectas");
        } else if (result.rol == 1) {
            window.open("/principal/admin", "_self");
            localStorage.setItem("Tipo", "1");
            localStorage.setItem("user", userName);
        } else if (result.rol == 2) {
            window.open("/principal/tutor", "_self");
            localStorage.setItem("Tipo", "2");
            localStorage.setItem("user", userName);
        } else if (result.rol == 3) {
            window.open("/principal/estudiante", "_self");
            localStorage.setItem("Tipo", "3");
            localStorage.setItem("user", userName);
        }
    };

    return (
        <div class="login-box">
            <h2>Inicio de Sesion
                Tutorias ECYS</h2>
            <form onSubmit={handleSubmit}>
                <div class="user-box">
                    <input type="text" id="userI" required value={userName} onChange={(e) => setUserName(e.target.value)} autoFocus />
                    <label>Username</label>
                </div>
                <div class="user-box">
                    <input type="password" id="passI" required aria-describedby="passwordHelpInline" value={passwordUser}
                        onChange={(e) => setPasswordUser(e.target.value)}
                        autoFocus />
                    <label>Password</label>
                </div>
                <div className="form-check form-switch text-left">
                    <input
                        className="form-check-input"
                        type="checkbox"
                        role="switch"
                        id="flexSwitchCheckDefault"
                        value={isChecked}
                        onChange={(e) => setIsChecked(!isChecked)}
                    />
                    <label
                        className="form-check-label"
                        htmlFor="flexSwitchCheckDefault"
                    >
                        <h3>Tutor</h3>
                    </label>
                </div>

                <button
                    className="w-100 btn btn-lg btn btn-outline-info"
                    type="submit"
                >
                    <span></span>
                    <span></span>
                    <span></span>
                    <span></span>
                    Iniciar Sesion
                </button>
                <br />
                <br />
                <h4>ESTRUCTURA DE DATOS 2023</h4>


            </form>
        </div>
    )
}

export default Login