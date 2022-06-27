$("#parar-de-seguir-usuario").on("click", pararDeSeguirUsuario)
$("#seguir-usuario").on("click", seguirUsuario)
$('#formulario-cadastro').on('submit', criarUsuario)
$("#editar-usuario").on('submit', editarUsuario)
$("#atualizar-senha").on('submit', atualizarSenha)
$("#deletar-usuario").on("click", deletarUsuario)


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

function pararDeSeguirUsuario(){
  const button = $(this)

  const usuarioID = button.data('usuario-id')
  button.prop('disabled', true)
  $.ajax({
    url: `/usuarios/${usuarioID}/parar-de-seguir`,
    method: "POST"
  }).done(function() {
    window.location = `/usuarios/${usuarioID}`
  }).fail(function() {
    Swal.fire("Ops...", "Erro ao parar de seguir o usuário!", "error")
    button.prop("disabled", false)
  })
}

function seguirUsuario() {
  const button = $(this)

  const usuarioID = button.data('usuario-id')
  button.prop('disabled', true)
  $.ajax({
    url: `/usuarios/${usuarioID}/seguir`,
    method: "POST"
  }).done(function() {
    window.location = `/usuarios/${usuarioID}`
  }).fail(function() {
    Swal.fire("Ops...", "Erro ao seguir o usuário!", "error")
    button.prop("disabled", false)
  })
}

function editarUsuario(event){
  event.preventDefault()

  var nome = $("#nome").val();
  var email = $("#email").val();
  var nick = $("#nick").val();

  $.ajax({
    url: "/editar-usuario",
    method: "PUT",
    data: {
      nome,
      email,
      nick
    }
  }).done(function() {
    Swal.fire(
    "Sucesso!", 
    "Usuário atualizado com sucesso!", 
    "success").then(function() {
      window.location = "/perfil"
    })
  }).fail(function(){
    Swal.fire(
      "Ops...", 
      "Erro ao atualizar o usuário!", 
      "error")
  })
}

function atualizarSenha(event){
  event.preventDefault()

  var senhaAtual = $("#senha-atual").val();
  var novaSenha = $("#nova-senha").val();
  var confirmarSenha = $("#confirmar-senha").val();

  if(novaSenha != confirmarSenha) {
    Swal.fire("Ops...", "As senhas não coincidem!", 'warning')
    return;
  }

  $.ajax({
    url: "/atualizar-senha",
    method: "POST",
    data: {
      atual: senhaAtual,
      nova: novaSenha
    }
  }).done(function(){
    Swal.fire("Sucesso", "A senha foi atualizada com sucesso!", 'success')
    .then(function(){
      window.location = "/perfil"
    })
  }).fail(function(){
    Swal.fire("Ops...","Erro ao atualizar a senha", "error")
  })
}

function deletarUsuario(){
  Swal.fire({
    title: "Atenção!",
    text: "Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversível!",
    icon: "warning",
    showCancelButton: true,
    cancelButtonText: "Cancelar"
  }).then(function(confirmacao){
    if (!confirmacao.value) return

    $.ajax({
      url: `/deletar-usuario`,
      method: "DELETE"
    }).done(function() {
      Swal.fire(
        "Sucesso", 
        "Seu usuário foi excluido com sucesso", 
        "success").then(function(){
          window.location = "/logout"
        })
    }).fail(function() {
      Swal.fire("Ops...", "Ocorreu um erro ao excluir o seu usuário", "error")
    });
  })
}