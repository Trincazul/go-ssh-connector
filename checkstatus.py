from paramiko import SSHClient
import paramiko
import psutil
from tqdm import tqdm
import time

class CheckStatus():

	def validatorup(ips,user,passw,command):

		ssh = paramiko.SSHClient()
		ssh.set_missing_host_key_policy(paramiko.AutoAddPolicy())
		ssh.connect(ips, username=user, password=passw)
		stdin, stdout, stderr = ssh.exec_command(command)
		print(stdout.read().decode())
		ssh.close()

if __name__ == '__main__':

	#configuração da conexão
	ips='0.0.0.0'
	user='root'
	passw='12345'

	print('Limpando cache')
	for i in tqdm(range(100)):
	    time.sleep(0.01)
	CheckStatus.validatorup(ips=ips,user=user,passw=passw,command='free -m')
	print('Aplicações Docker rodando')
	CheckStatus.validatorup(ips=ips,user=user,passw=passw,command='sudo docker ps')