<script>
  import { onMount } from "svelte";
  import { goto } from "$app/navigation";

  onMount(() => {
    setTimeout(() => {
      goto("/");
    }, 2500);
  });
</script>

<div class="modal_exito">
  <div class="modal_content">
    <h2>✅ Pedido realizado con éxito</h2>
    <p>Serás redirigido al inicio...</p>
  </div>
</div>

<style>
  :root {
    --background-light: #f9f9f9;
    --text-color: #333;
  }
  .modal_exito {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.6);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 9999;
    color: var(--text-color);
  }

  .modal_content {
    background: var(--background-light);
    padding: 30px;
    border-radius: 12px;
    text-align: center;
    font-size: 1.2rem;
  }
</style>
