<script lang="ts">
    import "../app.css"

    let { children } = $props()
</script>

<div class="h-screen flex flex-col">
    {@render children()}
</div>
