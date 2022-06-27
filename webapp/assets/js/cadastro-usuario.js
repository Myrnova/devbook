
$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(event) {
    event.preventDefault()

    var nome = $("#nome").val();
    var email = $("#email").val();
    var nick = $("#nick").val();
    var senha = $("#password").val();
    var confirmarPassword = $("#confirmar-password").val();

    if(senha != confirmarPassword) {
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome,
            email,
            nick,
            senha
        }
    }).done(function() {
        Swal.fire("Sucesso!", 
        "Usuário cadastrado com sucesso!", 
        "success").then(function() {
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email,
                    senha
                }
            }).done(function () {
                window.location.href = "/home";   
            }).fail(function(erro) {
                console.error(erro);
                Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
            })
        }) 
    }).fail(function(erro) {
        console.log(erro);
        Swal.fire("Ops...", "Erro ao cadastrar usuário!", "error");
    })
}