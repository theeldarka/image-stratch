<template>
    <div
        @dragenter.prevent="activateDropzone"
        @dragleave.prevent="deactivateDropzone"
        @dragover.prevent
        @drop.prevent="deactivateDropzone"
        :class="{ 'active-dropzone': active }"
        class="flex flex-col items-center justify-center w-full h-64 border-2 border-gray-300 border-dashed rounded-lg cursor-pointer bg-gray-50 transition-all duration-300 ease-in-out"
        @click="openFileBrowser"
    >
        <div class="pointer-events-none flex flex-col items-center justify-center pt-5 pb-6 transition-all duration-300 ease-in-out text-center">
            <svg aria-hidden="true" class="w-10 h-10 mb-3 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path></svg>

            <span v-if="active" class="mb-2 text-sm text-gray-500 dark:text-gray-400">
                {{ $t('dropzone.here') }}
            </span>
            <div class="flex flex-col" v-else>
                <span class="mb-2 text-sm text-gray-500 dark:text-gray-400"><span class="font-semibold">{{ $t('dropzone.click') }}</span> {{ $t('dropzone.dragAndDrop') }}</span>
                <span class="text-xs text-gray-500 dark:text-gray-400 text-center">{{ $t('dropzone.extensions') }}</span>
            </div>
        </div>

        <input type="file" accept="image/jpeg,image/png,image/bmp" ref="dropzoneFile" id="dropzoneFile" class="hidden" />
    </div>
</template>

<script>
export default {
    name: "DropZone",
    methods: {
        openFileBrowser() {
            this.$refs.dropzoneFile.click();
        },
        activateDropzone() {
            this.active = true;
        },
        deactivateDropzone() {
            this.active = false;
        },
    },
    data() {
        return {
            active: false,
        };
    },
};
</script>

<style>
    .active-dropzone {
        @apply text-white border-white bg-indigo-600;
    }

    .active-dropzone svg {
        @apply text-white;
    }

    .active-dropzone span {
        @apply text-white;
    }
</style>
