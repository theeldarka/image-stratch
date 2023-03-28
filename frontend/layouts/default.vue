<template>
    <div class="flex flex-col min-h-[90vh]">
        <main class="container mx-auto px-4" style="min-height: inherit">
            <LangSwitcher class="max-w-xs mx-auto mt-4" />
            <slot />
        </main>
        <footer class="mt-auto mb-2 px-4 flex justify-between md:justify-around flex-wrap">
            <span>Stretch Image &mdash; 2023</span>
            <span><a href="https://github.com/theeldarka/stretch-image" target="_blank">GitHub</a></span>
        </footer>
    </div>
</template>
