import chat1Img from './images/chat1.jpg'
import usuario1Img from './images/usuario1.jpg'

const u1 = {
    ID: 0,
    Nome: 'Yuri',
    FotoPerfil: '',
}
const u2 = {
    ID: 1,
    Nome: 'Henrique',
    FotoPerfil: usuario1Img,
}
const u3 = {
    ID: 3,
    Nome: 'Gabriel',
    FotoPerfil: '',
}

const m1 = {
    ID: 0,
    ChatID: 0,
    AutorID: 0,
    Conteudo: 'oi tudo bem?',
}
const m2 = {
    ID: 1,
    ChatID: 0,
    AutorID: 1,
    Conteudo: 'oiiiiie. Tudo bem e com vc?',
}
const m3 = {
    ID: 2,
    ChatID: 0,
    AutorID: 0,
    Conteudo: 'Estou ótimo também...só queria saber como vc estava msm!',
}
const m4 = {
    ID: 3,
    ChatID: 0,
    AutorID: 0,
    Conteudo: 'tchau bb',
}
const m5 = {
    ID: 4,
    ChatID: 0,
    AutorID: 1,
    Conteudo: 'xauuuuuuuu',
}

const c1 = {
    ID: 0,
    Nome: 'Chat 1',
    Usuarios: [
        u1, u2, u3
    ],
    Mensagens: [
        m1, m3
    ],
    FotoPerfil: chat1Img,
}

const c2 = {
    ID: 1,
    Nome: 'Henrique',
    Usuarios: [
        u1, u2
    ],
    Mensagens: [
        m1, m2, m3, m4, m5
    ],
    FotoPerfil: '',
}

const c3 = {
    ID: 2,
    Nome: 'Gabriel',
    Usuarios: [],
    Mensagens: [],
    FotoPerfil: '',
}

export const StubUsuarios = [
    u1, u2, u3
]

export const StubMensagens = [
    m1, m2, m3, m4
]

export const StubChats = [
    c1, c2, c3
]
