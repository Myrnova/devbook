
$('#formulario-login').on('submit', fazerLogin)


function fazerLogin(event) {
    event.preventDefault()

    var email = $("#email").val();
    var senha = $("#password").val();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email,
            senha
        }
    }).done(function() {
        alert("Login realizado com sucesso!");
        window.location.href = "/home";    
    }).fail(function(erro) {
        console.log(erro)
        alert("Usuário ou senha incorretos!");
    })
}