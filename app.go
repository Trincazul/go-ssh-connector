package main

import (
    "fmt"
    "log"
    "golang.org/x/crypto/ssh"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/mem"
    "io/ioutil"
)

func main() {

command := "cd Documents && ls -a"
user := "root"
// passw := "/home/root/.ssh/id_rsa"
passw := "yourpassword"
ip := "0.0.0.0"

// AuthRSA(user,passw,ip,command)

AuthUserPass(user,passw,ip,command)
}

// Função para acesso quando a ssh é feita com chave RSA
func AuthRSA(user string, privateKeyPath string, ips string, command string) {
    privateKey, err := ioutil.ReadFile(privateKeyPath)
    if err != nil {
        log.Fatalf("Failed to read private key: %s", err)
    }

    sshKey, err := ssh.ParsePrivateKey(privateKey)
    if err != nil {
        log.Fatalf("Failed to parse private key: %s", err)
    }

    sshConfig := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(sshKey),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    conn, err := ssh.Dial("tcp", ips+":22", sshConfig)
    if err != nil {
        log.Fatalf("Failed to dial: %s", err)
    }
    defer conn.Close()

    session, err := conn.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session: %s", err)
    }
    defer session.Close()

    // executa o comando 
    out, err := session.CombinedOutput(command)
    if err != nil {
        log.Fatalf("Failed to execute command: %s", err)
    }
    fmt.Println(string(out))
}

// Função para acesso quando a ssh é feita com senha
func AuthUserPass(user string, pass string, ips string, command string){

    sshConfig := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.Password(pass),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    // Conecta servidor remoto via SSH
    conn, err := ssh.Dial("tcp", ips+":22", sshConfig)
    if err != nil {
        log.Fatalf("Failed to dial: %s", err)
    }
    defer conn.Close()

    session, err := conn.NewSession()
    if err != nil {
        log.Fatalf("Failed to create session: %s", err)
    }
    defer session.Close()

    // executa o comando 
    out, err := session.CombinedOutput(command)
    if err != nil {
        log.Fatalf("Failed to execute command: %s", err)
    }
    fmt.Println(string(out))

    // verifica memória do computador
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        log.Fatalf("Failed to get CPU percent: %s", err)
    }
    memPercent, err := mem.VirtualMemory()
    if err != nil {
        log.Fatalf("Failed to get memory percent: %s", err)
    }
    fmt.Printf("Uso de CPU: %.2f%%\n", cpuPercent[0])
    fmt.Printf("Uso de memória: %.2f%%\n", memPercent.UsedPercent)
}
