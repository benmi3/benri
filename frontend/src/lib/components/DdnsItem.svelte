<article class="ddns-card">
	<h2>
		<p>
			Name:
			<slot name="name">
				<span class="missing">Unknown name</span>
			</slot>
		</p>
	</h2>
	<div class="domain">
		<p>
			Domain:
			<slot name="domain">
				<span class="missing">Unknown domain</span>
			</slot>
		</p>
	</div>
	{#if $$slots.a}
		<div class="a">
			<p>
				A:
				<slot name="a">
					<span class="missing">Unknown a record</span>
				</slot>
			</p>
		</div>
	{/if}
	{#if $$slots.aaaa}
		<div class="aaaa">
			<p>
				AAAA:
				<slot name="aaaa">
					<span class="missing">Unknown aaaa record</span>
				</slot>
			</p>
		</div>
	{/if}
	{#if $$slots.cname}
		<div class="cname">
			<p>
				CName:
				<slot name="cname">
					<span class="missing">Unknown cname</span>
				</slot>
			</p>
		</div>
	{/if}
</article>

<style>
	.ddns-card {
		width: 300px;
		border: 1px solid #aaa;
		border-radius: 2px;
		box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.1);
		padding: 1em;
	}
	p {
		padding: 0 0 0 0 0;
		margin: 0 0 0 0 0;
	}

	h2 {
		padding: 0 0 0.2em 0;
		margin: 0 0 1em 0;
		border-bottom: 1px solid #ff3e00;
	}

	.domain,
	.cname {
		padding: 0 0 0 1.5em;
		background: 0 0 no-repeat;
		background-size: 20px 20px;
		margin: 0 0 0.5em 0;
		line-height: 1.2;
	}

	.domain {
		background-image: url(/tutorial/icons/map-marker.svg);
	}
	.cname {
		background-image: url(/tutorial/icons/email.svg);
	}
	.missing {
		color: #999;
	}
</style>
