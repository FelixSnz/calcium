<template>
    <form @submit.prevent="SaveSearchReq()">
        <input class="bg-slate-600" type="text" placeholder="write keyword" v-model="SearchReq.keyword" />
        <input type="text" placeholder="write from date" v-model="SearchReq.from_date" />
        <input type="text" placeholder="write to date" v-model="SearchReq.to" />
        <button>Save</button>
    </form>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { SearchReq } from '../interfaces/SearchReq';

export default defineComponent({
    data() {
        return {
            SearchReq: {} as SearchReq
        }
    },
    methods: {
        SaveSearchReq() {
            console.log(this.SearchReq)
        }
    }
})
</script>