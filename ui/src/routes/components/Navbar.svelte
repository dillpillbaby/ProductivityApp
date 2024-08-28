<script>
    import {
        Navbar,
        NavBrand,
        NavLi,
        NavUl,
        NavHamburger,
    } from "flowbite-svelte";
</script>

<Navbar let:NavContainer color="gray">
    <NavContainer class="border px-5 py-1 rounded-lg bg-white dark:bg-gray-600">
        <NavBrand href="/">
            <img
                src="../../../logo.svg"
                class="me-3 h-24 sm:h-9"
                alt="Flowbite Logo"
            />
        </NavBrand>
        <NavHamburger />
        <NavUl>
            <NavLi href="/">Home</NavLi>
            <NavLi href="/tasks">Tasks</NavLi>
            <NavLi href="/reminders">Reminders</NavLi>
            <NavLi href="/about">About Me</NavLi>
        </NavUl>
    </NavContainer>
</Navbar>
