
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
        Swal.fire(
        "Sucesso",
        "Login realizado com sucesso!", 
        "success").then(function() {
            window.location.href = "/home";    
        });
    }).fail(function(erro) {
        console.log(erro)
        Swal.fire("Ops...","Usuário ou senha incorretos!", "error");
    })
}