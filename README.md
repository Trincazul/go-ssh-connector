# SSH Remote Access

Este é um código escrito em Go que permite acessar um servidor remoto via SSH. O código possui duas funções principais, que podem ser usadas para autenticar via chave RSA ou senha.
# Dependências

O código utiliza as seguintes bibliotecas externas:

   * golang.org/x/crypto/ssh: para realizar a conexão SSH
   * github.com/shirou/gopsutil/v3/cpu: para obter informações sobre o uso de CPU
   * github.com/shirou/gopsutil/v3/mem: para obter informações sobre o uso de memória

# Uso

O código possui duas funções principais: AuthRSA e AuthUserPass. A primeira é usada quando a autenticação é feita por meio de chave RSA, enquanto a segunda é usada quando a autenticação é feita por meio de senha.

Ambas as funções recebem os seguintes argumentos:

    `user`: o nome de usuário para autenticação
    `passw`: a senha ou caminho para a chave RSA (dependendo da função utilizada)
    `ip`: o endereço IP do servidor remoto
    `command`: o comando a ser executado no servidor remoto

Ao ser executado, o código imprime o resultado do comando no terminal, bem como informações sobre o uso de CPU e memória no servidor remoto.
