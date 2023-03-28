<template>
    <div>
        <label for="locale" class="block text-sm font-medium leading-6 text-gray-900">Language:</label>
        <select @change="(e) => {setLocale(e.target.value)}" id="locale" class="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900 ring-1 ring-inset ring-gray-300 focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6">
            <option v-for="locale in locales" :key="locale.code" :value="locale.code" :selected="locale.code === currentLocale">
                {{ locale.name }}
            </option>
        </select>
    </div>
</template>
<script setup>
const { locale: currentLocale, locales, setLocale } = useI18n()
</script>

<style>
    select {
        -webkit-appearance: none;
        appearance: none;

        background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' fill='none' viewBox='0 0 20 20'%3e%3cpath stroke='%236b7280' stroke-linecap='round' stroke-linejoin='round' stroke-width='1.5' d='M6 8l4 4 4-4'/%3e%3c/svg%3e");
        background-position: right 0.5rem center;
        background-repeat: no-repeat;
        background-size: 1.5em 1.5em;
        padding-right: 2.5rem;
        -webkit-print-color-adjust: exact;
        print-color-adjust: exact;
    }
</style>
