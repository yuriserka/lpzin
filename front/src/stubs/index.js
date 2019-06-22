import chat1Img from './images/chat1.jpg'
import henriqFoto from './images/usuario1.jpg'
import yuriFoto from './images/usuario2.jpg'
import gabrielFoto from './images/usuario3.jpg'

const u1 = {
    ID: 0,
    Nome: 'Yuri',
    FotoPerfil: yuriFoto,
    UltimaVez: '4 de jun de 2019',
}
const u2 = {
    ID: 1,
    Nome: 'Henrique',
    FotoPerfil: henriqFoto,
    UltimaVez: '47 minutos atrás',
}
const u3 = {
    ID: 2,
    Nome: 'Gabriel',
    FotoPerfil: gabrielFoto,
    UltimaVez: '7 horas atrás',
}
const u4 = {
    ID: 3,
    Nome: 'E',
    FotoPerfil: gabrielFoto,
    UltimaVez: '7 horas atrás',
}
const u5 = {
    ID: 4,
    Nome: 'A',
    FotoPerfil: gabrielFoto,
    UltimaVez: '7 horas atrás',
}
const u6 = {
    ID: 5,
    Nome: 'D',
    FotoPerfil: gabrielFoto,
    UltimaVez: '7 horas atrás',
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
const m6 = {
    ID: 5,
    ChatID: 0,
    AutorID: 2,
    Conteudo: 'eae galera bora fazer trabalho?',
}
const m7 = {
    ID: 6,
    ChatID: 0,
    AutorID: 2,
    Conteudo: 'semestre ta acabandooooo uhuuuuu',
}
const m8 = {
    ID: 7,
    ChatID: 0,
    AutorID: 1,
    Conteudo: 'cala boca corno',
}

const c1 = {
    ID: 0,
    Nome: 'Klub dos TeleZapCornos da UnB',
    Usuarios: [
        u1, u2, u3
    ],
    Mensagens: [
        m1, m2, m3, m4, m5, m6, m7, m8
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
    Usuarios: [
        u3, u1
    ],
    Mensagens: [],
    FotoPerfil: '',
}

export const StubUsuarios = [
    u1, u2, u3, u4, u5, u6
]

export const StubMensagens = [
    m1, m2, m3, m4, m5, m6, m7, m8
]

export const StubChats = [
    c1, c2, c3
]
