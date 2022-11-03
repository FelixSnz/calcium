
<template>
<div class="grow h-18 pl-6 pr-8 pt-2">

<form autocomplete="off" @submit.prevent="SaveSearchReq()" class="flex items-center">   
    <label for="voice-search" class="sr-only">Search</label>
    <div class="relative w-full p-2">
        <div class="flex absolute inset-y-0 left-2 items-center pl-3 pointer-events-none">
            <svg class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clip-rule="evenodd"></path></svg>
        </div>
        <input autocomplete="on" type="text" id="voice-search" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search by keyword..." required v-model="SearchReq.keyword">
        <button type="button" class="flex absolute inset-y-0 right-2 items-center pr-3">
            <svg class="w-4 h-4 text-gray-500 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M7 4a3 3 0 016 0v4a3 3 0 11-6 0V4zm4 10.93A7.001 7.001 0 0017 8a1 1 0 10-2 0A5 5 0 015 8a1 1 0 00-2 0 7.001 7.001 0 006 6.93V17H6a1 1 0 100 2h8a1 1 0 100-2h-3v-2.07z" clip-rule="evenodd"></path></svg>
        </button>
        
    </div>
    

    <div date-rangepicker datepicker-autohide class="flex w-200">
        <div class="relative grow">
            <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"></path></svg>
            </div>
            <input ref="startdate" name="start" type="text" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-200 pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Select date start">
        </div>
        <span class="mx-4 text-gray-500">to</span>
        <div class="relative grow">
            <div class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
                <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd"></path></svg>
            </div>
            <input ref="enddate" name="end" type="text" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-200 pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Select date end">
        </div>
        </div>
        <button type="submit" class="inline-flex items-center py-2.5 px-3 ml-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"><svg class="mr-2 -ml-1 w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>Search</button>
        
</form>

<div class="flex-grow">
            <div class="flex px-1 py-3">
<div class="overflow-auto shadow-md sm:rounded-lg">
                    <div class="inline-block min-w-full max-w-4xl h-96 max-h-full align-middle">
                <div class="overflow-auto ">
                    <table class="table-auto grow text-sm text-left text-gray-500 dark:text-gray-400">
                        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                            <tr>
                            <th scope="col" class="px-4 py-3">Origin</th>
                            <th scope="col" class="px-6 py-3">SubFolder</th>
                            <th scope="col" class="px-10 py-3">Date</th>
                            <th scope="col" class="px-8 py-3">From</th>
                            <th scope="col" class="px-6 py-3">To</th>
                            <th scope="col" class="px-6 py-3">Cc</th>
                            </tr>
                        </thead>
                        <tbody v-for="(mail, index) in emails" :key="index">
                            <tr :class="[(active_row == index ? 'active_row' : '') ,'onHover']" v-on:click="HandleEmailClick(mail, index)"  class=" bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                            <td class="px-6 py-4">{{mail.origin}}</td>
                            <td class="px-6 py-4">{{mail.subfolder}}</td>
                            <td class="px-6 py-4">{{new Date(mail.date)}}</td>
                            <td class="px-6 py-4">{{mail.from}}</td>
                            <td v-bind:title="mail.to" class="px-6 py-4">...</td>
                            <td v-bind:title="mail.cc" class="px-6 py-4">...</td>
                            
                            </tr>

                        </tbody>
                        </table>
                    </div>
                </div>
                </div>

                <div class="flex grow pl-5">
                <div>
    
        <div rows="1" class="px-5 grow block w-100 min-w-full text-sm text-gray-900 bg-gray-50 rounded-t-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="...">
               <b>Subject: </b> {{selectedEmail.subject}}
        </div>
        <div v-html="selectedEmail.content" type="text" id="message" rows="15" class="overflow-auto max-w-md max-h-96 px-5 grow block w-100 min-w-full text-sm text-gray-900 bg-gray-50 rounded-b-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="...">
        </div>

        </div>
        </div>
        <!-- <div v-html="content" type="text" id="message" rows="15" class="overflow-auto max-w-md max-h-96 px-5 grow block w-100 min-w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="...">
            
        </div> -->
    
    </div>
        </div>

</div>




</template>

<script lang="ts">
import { defineComponent } from 'vue';
import type { SearchReq } from '@/interfaces/SearchReq';
import type {EmailData} from '@/interfaces/EmailData'
import SearchEmail from '@/services/SearchService';

//or in file components


export default defineComponent({
    data() {
        return {
            SearchReq: {} as SearchReq,
            emails: [] as EmailData[],
            selectedEmail: {} as EmailData,
            message:  "" as string,
            rowSelected : false,
            active_row: 0
        }
    },
    methods: {
        async SaveSearchReq() {


            console.log(this.SearchReq.keyword)


            const from_date = new Date((this.$refs.startdate as any).value).toISOString();
            const to_date = new Date((this.$refs.enddate as any).value).toISOString();
            this.SearchReq.from = from_date
            this.SearchReq.to = to_date
            const res = await SearchEmail(this.SearchReq)

            this.emails = res.data

        },
        HandleEmailClick (email: EmailData, idx: number) {
            console.log(email.content)
            var searchMask = this.SearchReq.keyword;
            var regEx = new RegExp(searchMask, "ig");

            this.selectedEmail = email


            var replaceMask = this.SearchReq.keyword.toUpperCase().bold();
            this.selectedEmail.content = email.content.replace('\n', '<br\>').replace('\t', '<pre\>').replace(regEx, replaceMask);
            this.active_row = idx

            },


        // MouseOver: function (str:string) {
        //         this.replaced = str.replace(",", "\n");    
        //     },
        myFunction (str:string) {
            this.message = str.replace("ItSolutionStuff.com", "HDTuto.com");   
            }
    },
})
</script> -->


    
    