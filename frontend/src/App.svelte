<script>
  import { onMount } from "svelte";
  import { TodosService } from "./oto.gen.js";
  import Todo from "./Todo.svelte";
  const todoService = new TodosService();
  let todos = [];
  let error = null;
  let task = "";

  onMount(async () => {
    todoService
      .all({})
      .then(response => (todos = response.todos))
      .catch(error => console.error);
  });

  const handleSubmit = event => {
    todoService
      .create({
        todo: {
          task: task,
          done: false
        }
      })
      .then(response => {
        todos = [response.todo, ...todos];
        task = "";
        error = null;
      })
      .catch(err => {
        error = err;
      });
  };

  function mark(todo) {
    remove(todo);
    todos = todos.concat(todo);
  }

  function remove(todo) {
    todos = todos.filter(t => t.id !== todo.id);
  }
</script>

<style>
  header {
    padding: 1rem 0;
  }
  main {
    text-align: center;
    padding: 1em;
    max-width: 480px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }

  ul {
    display: flex;
    flex-direction: column;
    justify-content: center;
    margin: 0;
    padding: 1rem 0;
    width: 100%;
  }

  h1,
  h2,
  h3 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
    margin: 0;
  }

  h2 {
    font-size: 3em;
  }
  h3 {
    font-size: 2em;
  }

  form {
    width: 100%;
  }

  input {
    width: 100%;
    font-size: 1.5rem;
    padding: 0.5rem;
    margin: 0;
  }

  p.danger {
    color: #c00;
    font-weight: bold;
    font-size: 1.2rem;
    padding: 1rem 0.5rem;
    text-align: left;
    margin: 0;
  }
</style>

<main>
  <header>
    <h1>TODOS!</h1>
  </header>

  <form on:submit|preventDefault={handleSubmit}>
    <input
      type="text"
      name="task"
      placeholder="What needs to be done?"
      bind:value={task} />
    {#if error}
      <p class="danger">{error}</p>
    {/if}
  </form>
  <ul>
    {#each todos.filter(t => !t.completed) as todo (todo.id)}
      <Todo {...todo} {mark} {remove} />
    {/each}
  </ul>

  <h3>Done!</h3>
  <ul>
    {#each todos.filter(t => t.completed) as todo (todo.id)}
      <Todo {...todo} {mark} {remove} />
    {/each}
  </ul>
</main>
