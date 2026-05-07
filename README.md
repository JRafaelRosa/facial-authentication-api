# facial-authentication-api

Sistema híbrido de autenticação biométrica baseado em reconhecimento facial, composto por um módulo de visão computacional em Python e uma API backend em Go responsável por validação, regras de negócio e persistência de logs.

---

## 🧠 Visão Geral do Sistema

O sistema simula um controle de acesso biométrico dividido em dois componentes principais:

- **Python (Visão Computacional)**: Captura e reconhece faces em tempo real.
- **Go (Backend)**: Valida os dados recebidos, aplica regras de negócio e registra logs no banco de dados.

---

## 🔄 Fluxo do Sistema

1. **Captura**: Câmera / Imagem  
2. **Processamento**: Python (OpenCV + LBPH)  
3. **Identificação**: Geração de ID e Accuracy (Confiança)  
4. **Comunicação**: Requisição HTTP POST  
5. **API Go**: Recebimento no endpoint `/registrar`  
6. **Negócio**: Validação de regras e permissões  
7. **Persistência**: Registro de logs no MySQL  

---

## ⚙️ Regras de Negócio

- **Acesso Autorizado**: `accuracy >= threshold` (ex: 80%).
- **Acesso Negado**: `accuracy < threshold`.

> Todos os eventos (tentativas de acesso) são registrados no banco de dados, independentemente do resultado da autenticação.

---

## 🧩 Arquitetura

### Backend (Go)

Estruturado seguindo princípios de modularidade para facilitar a manutenção e escalabilidade:
- **handler**: Camada de entrada HTTP e manipulação de requests.
- **service**: Implementação das regras de negócio.
- **storage**: Interface de comunicação e persistência no banco.
- **model**: Definição das entidades do sistema.

**Tecnologias**: Go (Golang), MySQL.

---

### Visão Computacional (Python)

Responsável pelo processamento de imagem na "borda":
- Detecção facial (Haar Cascades).
- Reconhecimento com OpenCV (LBPH).
- Geração de score de confiança e normalização de imagem.

**Link do Repositório**: [face-recognition-opencv](https://github.com/JRafaelRosa/face-recognition-opencv)

**Tecnologias**: Python 3, OpenCV, NumPy.

---

## 📡 Integração entre os sistemas

A comunicação entre os módulos ocorre via JSON sobre HTTP.

**Exemplo de Payload:**
```json
{
  "name": "Joao",
  "email": "jrafael@teste.com",
  "position": "developer",
  "accuracy": 87.5
}
```

---
Desenvolvido por Joao Rafael dos Santos da Rosa. Estudante de Engenharia de Computação (UEPG) e Desenvolvedor Full Stack.
