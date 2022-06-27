$("#nova-publicacao").on('submit', criarPublicacao)


$(document).on('click', '.curtir-publicacao', curtirPublicacao) //mudamos para document porque houve atualização das classes do componente, então é necessário atrelhar ao document e nao a classe em si
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao)
$("#atualizar-publicacao").on('click', atualizarPublicacao)
$(".deletar-publicacao").on('click', deletarPublicacao)


function criarPublicacao (event) {
  event.preventDefault();
  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $("#titulo").val(),
      conteudo: $("#conteudo").val()
    }
  }).done(function () {
    window.location = "/home"
  }).fail(function () {
    Swal.fire("Erro", "Erro ao criar publicação", "error");
  })
}

function curtirPublicacao(event) {
  const elementoClicado = $(event.target)
  const publicacaoId = elementoClicado.closest('.jumbotron').data('publicacao-id')

  elementoClicado.prop('disabled', true)
  $.ajax({
    url: `/publicacoes/${publicacaoId}/curtir`,
    method: "POST"
  }).done(function () {
    const contadorDeCurtidas = elementoClicado.next('span')
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())
    contadorDeCurtidas.text(quantidadeDeCurtidas + 1)

    elementoClicado.addClass("descurtir-publicacao")
    .addClass("text-danger")
    .removeClass("curtir-publicacao")

  }).fail(function () {
    Swal.fire("Ops...", "Erro ao curtir publicacao!", "error")
  }).always(function () { //always sempre roda independente de ter dado certo ou errado a requisição
    elementoClicado.prop('disabled', false)
  })
  
}

function descurtirPublicacao(event) {
  const elementoClicado = $(event.target)
  const publicacaoId = elementoClicado.closest('.jumbotron').data('publicacao-id')

  elementoClicado.prop('disabled', true)
  $.ajax({
    url: `/publicacoes/${publicacaoId}/descurtir`,
    method: "POST"
  }).done(function () {
    const contadorDeCurtidas = elementoClicado.next('span')
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())
    contadorDeCurtidas.text(quantidadeDeCurtidas - 1)

    elementoClicado.addClass("curtir-publicacao")
    .removeClass("text-danger")
    .removeClass("descurtir-publicacao")

  }).fail(function () {
    Swal.fire("Ops...", "Erro ao descurtir publicacao!", "error")
  }).always(function () { //always sempre roda independente de ter dado certo ou errado a requisição
    elementoClicado.prop('disabled', false)
  })
  
}

function atualizarPublicacao() {
  const button = $(this)
  button.prop('disabled', true);

  const publicacaoId = button.data("publicacao-id");

  $.ajax({
    url: `/publicacoes/${publicacaoId}`,
    method: "PUT",
    data: {
      titulo: $("#titulo").val(),
      conteudo: $("#conteudo").val(),
    }
  }).done(function() {
    Swal.fire(
    "Sucesso!", 
    "Publicação editada com sucesso!",
    "success")
    .then(function(){
      window.location = "/home"
    });
  }).fail(function(erro) {
    Swal.fire("Erro!", "Erro ao editar a publicação!", "error");
  }).always(function() {
    console.log(button)
    button.prop("disabled", false)
  })

}

function deletarPublicacao(event) {
  event.preventDefault()

  Swal.fire({
    title: "Atenção",
    text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
    showCancelButton: true,
    cancelButtonText: "Cancelar",
    icon: "warning"
  }).then(function(confirmacao) {
    if (!confirmacao.value) return;

  const elementoClicado = $(event.target)
  const publicacaoDiv = elementoClicado.closest('.jumbotron')
  const publicacaoId = publicacaoDiv.data('publicacao-id')

  elementoClicado.prop('disabled', true)
  $.ajax({
    url: `/publicacoes/${publicacaoId}`,
    method: "DELETE"
  }).done(function () {
    publicacaoDiv.fadeOut("slow", function() {
      $(this).remove();
    })
  }).fail(function (error) {
    console.error(error)
    Swal.fire("Ops...", "Erro ao excluir publicacao!", "error");
  }).always(function () { //always sempre roda independente de ter dado certo ou errado a requisição
    elementoClicado.prop('disabled', false)
  })
  })

  
}