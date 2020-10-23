# Curso: Desenvolvimento de Aplicações Modernas e Escaláveis com Microsserviços

## Modulo: DevOps - Kubernetes e hpa

### Desafio: Kubernetes e HPA (Horizontal Pod Autoscaler)

1. Aplicação em Golang
   - Criar um algoritimo quee execute um looping somando a raiz quadrada.
   - Criar os arquivos de configuração para o deployment, service do kubernetes.
   - Cada réplica deverá consumir no mínimo 50m e no máximo 100m.

2. Criar um arquivo de configuração para o hpa
   - O processo de escala inicia quando a CPU passar de 15%
   - Quantidade mínima de pods: 1
   - Quantidade máxima de pods: 6

3. Crie um POD e faça requisições através de um looping infinito e verifique se o autoscaler está funcionando corretamente.

---

***Aplicação em Golang***

- *Execução a aplicação*
```
# Rodando o main
> go run src/raiz/main.go

# Gerando o build
> go build src/raiz/main.go
# Executando
> ./main
```

Acesse o link: http://localhost:8000/, aparecerá no browser "Raiz: 1.0. Code.education Rocks!"

- *Rodando os testes*
```
> go test -v src/raiz/main_test.go src/raiz/main.go
=== RUN   TestMainSuccess
    main_test.go:14: Function sum success
--- PASS: TestMainSuccess (0.01s)
PASS
ok  	command-line-arguments	1.081s
``` 

- *Gerando uma imagem Docker*
```
# Gerando uma imagem docker
> docker build -t leticiapillar/go-hpa .

# Executando a imagem docker
> docker run -d -p 8000:8000 leticiapillar/go-hpa
```

Acesse o link: http://localhost:8000/, aparecerá no browser "Raiz: 1.0. Code.education Rocks!"

Imagem no Docker Hub: https://hub.docker.com/r/leticiapillar/go-hpa

---

***Arquivos de configurações para o Kubernetes***

- Deployment: `go-deployment.yaml`
- Service: `go-service.yaml`
- HPA: `go-hpa.yaml`

```
# Aplicar as configurações de deployment
> kubectl apply -f go-deployment.yml

# Aplicar as configurações de service
> kubectl apply -f go-service.yml

# Testas a execução pelo minikube
> minikube service go-hpa

# Aplicar as configurações do hpa
> kubectl apply -f go-hpa.yml

```

***POD para testar um loop de requisições***

- Criar um POD `loader` com a imagem busybox
```
> kubectl run -it loader --image=busybox /bin/sh
If you don't see a command prompt, try pressing enter.
/ # wget -q -O- http://go-hpa.default.svc.cluster.local;
/ # while true; do wget -q -O- http://go-hpa.default.svc.cluster.local; done;
Code.education Rocks!Raiz: 1.0. Code.education Rocks!Raiz: 1.0. Code.education Rocks!Raiz: 1.0. Code.education Rocks!Raiz: 1.0. Code.education Rocks!Raiz: 1.0. Code.education Rocks!Raiz: 1.0. Code.education Rocks!^C
/ # exit
```

- Verificando o autoscaler
```
# Execução única
> kubectl get hpa
NAME     REFERENCE           TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
go-hpa   Deployment/go-hpa   0%/15%    1         5         1          32m

# Monitora o hpa
> watch kubectl get hpa
```

- Referências:
  * [Horizontal Pod Autoscaler Walkthrough](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/)