<script lang=ts>
    import Navbar from "../components/Navbar.svelte";
    import { Card } from "flowbite-svelte";
    
    let description = '';
    let mood = 0;

    async function handleSubmit(event: Event) {
        // Prevent the default form submission behavior
        event.preventDefault();

        // Create the object to send to the backend
        const journalEntry = {
            description,
            'MoodID': mood
        };

        try {
            // Make a POST request to your endpoint
            const response = await fetch('http://localhost:8080/journalentries', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(journalEntry),
            });
            console.log(JSON.stringify(journalEntry))
            
            if (!response.ok) {
                throw new Error('Failed to create journal entry');
            }
            
            // Handle the response (e.g., clear the form, show a success message)
            const result = await response.json();
            console.log('Journal entry created:', result);

            // Reset form fields if needed
            description = '';
            mood = 0;
        } catch (error) {
            console.error('Error:', error);
        }
    }
</script>

<Navbar></Navbar>
<form class="flex flex-col" on:submit={handleSubmit}>
    <input type="text" class="w-48" placeholder="Description" bind:value="{description}" />
    <input type="number" class="w-48" placeholder="Mood" bind:value="{mood}" />
    <button type="submit" class="w-48 h-12 border"
        ><span>Create Journal Entry</span></button
    >
</form>
<div class="flex w-full h-full justify-center mt-4 flex-wrap mx-72">
    <Card class="mx-4 max-w-xs">
        <h5
            class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
        >
            Title
        </h5>
        <p class="font-normal text-gray-700 dark:text-gray-400 leading-tight">
            body
        </p>
    </Card>
    <Card class="mx-4 max-w-xs">
        <h5
            class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
        >
            Title
        </h5>
        <p class="font-normal text-gray-700 dark:text-gray-400 leading-tight">
            body
        </p>
    </Card>
    <Card class="mx-4 max-w-xs">
        <h5
            class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
        >
            Title
        </h5>
        <p class="font-normal text-gray-700 dark:text-gray-400 leading-tight">
            body
        </p>
    </Card>
</div>
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
        margin-top: 150px;
    }
</style>
