<template>
    <form @submit="(e) => {e.preventDefault()}"
          class="container mx-auto flex flex-row gap-4 justify-center"
          style="min-height: inherit">
        <div class="w-[350px] flex flex-col gap-4 justify-center my-auto">
            <fieldset>
                <legend class="block text-sm font-medium leading-6 text-gray-900">{{ $t('form.output_size') }}</legend>
                <div class="mt-2 -space-y-px rounded-md bg-white shadow-sm">
                    <div class="flex -space-x-px">
                        <div class="w-1/2 min-w-0 flex-1">
                            <label for="width" class="sr-only">{{ $t('form.width') }}</label>
                            <input v-model="width"
                                   type="text"
                                   name="width"
                                   id="width"
                                   class="text-center relative block w-full rounded-l-md border-0 bg-transparent py-1.5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                   placeholder="1920"/>
                        </div>
                        <div class="min-w-0 flex-1">
                            <label for="height" class="sr-only">{{ $t('form.height') }}</label>
                            <input v-model="height"
                                   type="text"
                                   name="height"
                                   id="height"
                                   class="text-center relative block w-full rounded-r-md border-0 bg-transparent py-1.5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:z-10 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                                   placeholder="1080"/>
                        </div>
                    </div>
                </div>
            </fieldset>

            <fieldset class="flex align-baseline justify-between">
                <span class="font-medium text-sm">{{ $t('form.black_background') }}</span>
                <Switch v-model="wantBlackBackground"
                        :class="[wantBlackBackground ? 'bg-indigo-600' : 'bg-gray-200', 'relative inline-flex h-6 w-11 flex-shrink-0 cursor-pointer rounded-full border-2 border-transparent transition-colors duration-200 ease-in-out focus:outline-none focus:ring-2 focus:ring-indigo-600 focus:ring-offset-2']">
                    <span aria-hidden="true"
                          :class="[wantBlackBackground ? 'translate-x-5' : 'translate-x-0', 'pointer-events-none inline-block h-5 w-5 transform rounded-full bg-white shadow ring-0 transition duration-200 ease-in-out']"/>
                </Switch>
                <input type="hidden" id="black_background" name="black_background" :value="wantBlackBackground">
            </fieldset>

            <DropZone @drop.prevent="fileDropped" @change="fileSelected"/>
        </div>
    </form>
</template>

<script>
import {Switch} from '@headlessui/vue'
import DropZone from "@/components/DropZone.vue";
import { useToast } from 'vue-toastification/dist/index.mjs'

export default {
    components: {
        Switch,
        DropZone,
    },
    data() {
        return {
            wantBlackBackground: true,
            width: 1920,
            height: 1080,
            dropzoneFile: null,
        }
    },
    // computed: {
    //     color() {
    //         if (this.wantBlackBackground) {
    //             return { red: 0, green: 0, blue: 0 }
    //         }
    //
    //         return { red: 255, green: 255, blue: 255 }
    //     },
    // },
    methods: {
        fileDropped(e) {
            this.dropzoneFile = e.dataTransfer.files[0];

            this.sendFile(this.dropzoneFile)
        },
        fileSelected() {
            this.dropzoneFile = document.querySelector("#dropzoneFile").files[0];

            this.sendFile(this.dropzoneFile)
        },
        sendFile(selectedFile) {
            const formData = new FormData()

            formData.append('image', selectedFile)

            formData.append('width', this.width)
            formData.append('height', this.height)
            formData.append('black_background', this.wantBlackBackground)

            // Object.keys(this.color).forEach(key => {
            //     formData.append('color[' + key + ']', this.color[key])
            // })
            //
            // Object.keys(this.color).forEach(key => {
            //     formData.append('color.' + key, this.color[key])
            // })
            // formData.append('color', JSON.stringify(this.color))

            // Submit the form using the FormData object
            fetch('/process-image', {
                method: 'POST',
                body: formData
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error(response.statusText);
                    }

                    return response.blob();
                })
                .then(this.downloadFile)
                .catch(error => {
                    console.error('Error:', error);

                    useToast().error(this.$t('status.error') + ': ' + error.message);
                });
        },
        downloadFile(blob) {
            const url = URL.createObjectURL(blob);
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute('download', 'image.jpg');
            document.body.appendChild(link);
            link.click();

            useToast().success(this.$t('status.success'));
        },
    },
}
</script>
