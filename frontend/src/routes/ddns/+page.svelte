<script>
	import Loading from '../../lib/components/Loader.svelte';
	import DdnsRecord from '../../lib/components/DdnsItem.svelte';
</script>

<svelte:head>
	<title>Ddns</title>
	<meta name="description" content="About this app" />
</svelte:head>

<div class="text-column">
	<h1>Ddns manager</h1>

	<p>
		This is a <a href="https://kit.svelte.dev">SvelteKit</a> app. You can make your own by typing the
		following into your command line and following the prompts:
	</p>

	<pre>bun create svelte@latest</pre>

	<p>
		The page you're looking at is purely static HTML, with no client-side interactivity needed.
		Because of that, we don't need to load any JavaScript. Try viewing the page's source, or opening
		the devtools network panel and reloading.
	</p>

	<p>
		The <a href="/sverdle">Sverdle</a> page illustrates SvelteKit's data loading and form handling. Try
		using it with JavaScript disabled!
	</p>

	<Loading />
	<DdnsRecord>
		<span slot="name"> auth </span>
		<span slot="domain"> test.example </span>
		<span slot="cname"> test.example </span>
	</DdnsRecord>
</div>
