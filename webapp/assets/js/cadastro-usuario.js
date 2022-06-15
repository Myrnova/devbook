
$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(event) {
    event.preventDefault()

    var nome = $("#nome").val();
    var email = $("#email").val();
    var nick = $("#nick").val();
    var password = $("#password").val();
    var confirmarPassword = $("#confirmar-password").val();

    if(password != confirmarPassword) {
        alert("As senhas não coincidem!");
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome,
            email,
            nick,
            password
        }
    })
}