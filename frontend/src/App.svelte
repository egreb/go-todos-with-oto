<script>
  import { onMount } from "svelte";
  import { TodosService } from "./oto.gen.js";
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
</script>

<style>
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

  li {
    border: 1px solid #999;
    box-shadow: 1 1 1 1 #333;
    width: 100%;
    border-radius: 5px;
    margin-top: 5px;
    list-style: none;
    padding: 0.5rem;
    font-size: 1.5rem;
  }

  li:first-child {
    margin-top: 0;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
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
  <h1>TODOS!</h1>

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
    {#each todos as todo}
      <li>{todo.task}</li>
    {/each}
  </ul>
</main>
