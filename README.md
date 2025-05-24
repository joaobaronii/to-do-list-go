# Go To-Do List CRUD

Este é um aplicativo simples de linha de comando (CLI) para gerenciar uma lista de tarefas (To-Do List). Ele permite adicionar, listar, marcar como concluídas e excluir tarefas. Os dados são armazenados em um banco de dados PostgreSQL.

## Funcionalidades

* **Adicionar Tarefa**: Permite ao usuário adicionar uma nova tarefa à lista.
* **Listar Tarefas**:
    * Listar todas as tarefas.
    * Listar apenas tarefas concluídas.
    * Listar apenas tarefas não concluídas.
* **Marcar Tarefa como Concluída**: Permite ao usuário marcar uma tarefa existente como concluída.
* **Excluir Tarefa**:
    * Excluir todas as tarefas.
    * Excluir todas as tarefas concluídas.
    * Excluir uma tarefa específica pelo nome.
* **Sair**: Encerra a aplicação.

## Tecnologias Utilizadas

* **Go**: Linguagem de programação principal.
* **PostgreSQL**: Sistema de gerenciamento de banco de dados relacional.
* **pgx**: Driver Go para PostgreSQL.
* **Viper**: Para gerenciamento de configuração (leitura de variáveis de ambiente de um arquivo `.env`).
* **UUID**: Para gerar IDs únicos para as tarefas.

## Lógica Principal dos Arquivos

* `main.go`: Contém a função `main` que inicializa a conexão com o banco de dados, cria a tabela de tarefas (se não existir) e exibe o menu principal para interação com o usuário. Também contém as funções `menuListTasks`, `taskDone`, e `menuDeleteTask` que gerenciam os respectivos submenus e interações.
* `config.go` (`configs/`): Define a struct `config` e a função `LoadConfig` para carregar as configurações do banco de dados a partir de um arquivo `.env` usando a biblioteca Viper. A função `GetPostgresURL` formata a URL de conexão do PostgreSQL.
* `taskDB.go` (`internal/database/`): Responsável por todas as operações diretas com o banco de dados relacionadas às tarefas, como criar a tabela, inserir, selecionar (todas, por status), marcar como concluída e excluir tarefas.
* `task.go` (`internal/entity/`): Define a struct `Task` que representa uma tarefa com ID, Nome e Status. Inclui um construtor `NewTask` que inicializa uma nova tarefa com um UUID.
* `taskService.go` (`internal/models/`): Contém as operações de tarefa. A função `AddTask` lida com a entrada do usuário para o nome da nova tarefa e chama a função de inserção no banco de dados. `ListTasks` formata e exibe a lista de tarefas. `DeleteTaskByName` lida com a entrada do usuário para excluir uma tarefa específica.

