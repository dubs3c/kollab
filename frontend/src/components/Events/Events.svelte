<script lang="ts">
	import { GetEvents } from "../../actions/DefaultPathAction.svelte";
	import type { LogEvent } from "src/models";
	import { onInterval } from "./utils";


    export let pathId = ""
    let events: LogEvent[] = []

    async function populateEvents() {
        events = await GetEvents(pathId);
    }

    onInterval(populateEvents, 3000);
</script>

<div class="card">
    <div class="card-body">
        <div id="eventlog" style="width: 100%; height: 200px; overflow:scroll">
            {#each events as event }
                <span style="color:blueviolet">{event.Path}</span> <span style="color:tomato">{event.IP}</span> {#each Object.entries(event.RequestHeaders) as [key, value]}
                    <span style="color:chocolate">[{key} {value}]</span>
                {/each}
                <br />
            {/each}
        </div>
    </div>
</div>