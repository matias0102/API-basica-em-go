package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	//CRUD DE USUARIOS
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuarios,
		RequerAltenticacao: false,
	},

	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsusarios,
		RequerAltenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsusario,
		RequerAltenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuarios,
		RequerAltenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuarios,
		RequerAltenticacao: true,
	},
	//-------------- Seguidores -----------------

	{
		URI:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAltenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}/not-follow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.NotFollow,
		RequerAltenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAltenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguindo,
		RequerAltenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}/atualizar-Senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAltenticacao: true,
	},
}
