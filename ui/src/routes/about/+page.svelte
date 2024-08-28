<script>
   import Navbar from "../components/Navbar.svelte";
</script>

<Navbar>
</Navbar>
<span>About</span>

<!-- <div id="topBar" style="display: flex;">
    <img
        src="..\logo.png"
        alt="logo"
        style="display: flex; width: 100px; border-bottom: 2cap; border-bottom: gray; position:fixed;"
    />
    <div id="navigationBar" class="navigationBar">
        {#if navigationBarOpen}
            <div
                id="navigationBarButtons"
                role="banner"
                on:mouseleave={closeNavigationBar}
            >
                <button class="navigationBarButton"></button>
                <button class="navigationBarButton"></button>
                <button class="navigationBarButton"></button>
            </div>
        {:else}
            <svg
                on:mouseenter={openNavigationBar}
                role="contentinfo"
                width="24"
                height="24"
                xmlns="http://www.w3.org/2000/svg"
                fill-rule="evenodd"
                clip-rule="evenodd"
                ><path
                    d="M11 21.883l-6.235-7.527-.765.644 7.521 9 7.479-9-.764-.645-6.236 7.529v-21.884h-1v21.883z"
                /></svg
            >
        {/if}
    </div>
</div> -->
<p class="footer">Made by Dylan Hawkins for the "Project Jam 2024"</p>

<style>
    .footer {
        text-align: center;
        font-size: 1em;
        padding-top: 150px;
    }
</style>
