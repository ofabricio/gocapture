package main

const dashboardHTML = `
<!DOCTYPE html>
<html ng-app="app">
<head>
	<meta charset="utf-8">
	<link rel="icon" href="data:;base64,iVBORw0KGgo=">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<script src="https://cdn.jsdelivr.net/npm/vue@2.6.0"></script>
	<link href="https://fonts.googleapis.com/css?family=Inconsolata:400,700" rel="stylesheet">
	<title>Dashboard</title>
	<style>

	[v-cloak] { display: none !important }

	:root {
		--bg: #282c34;
		--list-item-bg: #2c313a;
		--list-item-fg: #abb2bf;
		--list-item-sel-bg: hsl(219, 22%, 25%);
		--req-res-bg: #2c313a;
		--req-res-fg: #abb2bf;
		--links: #55b5c1;
		--method-get: #98c379;
		--method-post: #c678dd;
		--method-put: #d19a66;
		--method-patch: #a7afbc;
		--method-delete: #e06c75;
		--status-ok: #98c379;
		--status-warn: #d19a66;
		--status-error: #e06c75;
		--btn-bg: var(--list-item-bg);
		--btn-hover: var(--list-item-sel-bg);
		--disabled: hsl(187, 5%, 50%);
	}

	* { padding: 0; margin: 0; box-sizing: border-box }

	div { position: relative }

	html, body, .dashboard {
		height: 100%;
		font-family: 'Inconsolata', monospace;
		font-size: 1em;
		font-weight: 400;
		background: var(--bg);
	}

	.dashboard {
		display: grid;
		grid-template-columns: .6fr 1fr 1fr;
		grid-template-rows: 1fr;
		grid-gap: .5rem;
	}

	.list, .req, .res {
		display: grid;
		grid-template-rows: auto 1fr;
		grid-gap: .5rem;
	}

	body { padding: .5rem; }

	.list, .req, .res {
		overflow: auto;
	}

	.list-inner, .req-inner, .res-inner {
		overflow-x: hidden;
		overflow-y: auto;
	}

	.req-inner, .res-inner {
		background: var(--req-res-bg);
	}

	.req, .res {
		color: var(--req-res-fg);
	}

	.list-inner {
		display: grid;
		grid-template-rows: auto;
		grid-gap: .5rem;
		align-content: start;
	}

	.list-item {
		display: grid;
		grid-template-columns: auto 1fr auto auto;
		grid-gap: .5rem;
		font-size: 1.2em;
		padding: 1rem;
		background: var(--list-item-bg);
		color: var(--list-item-fg);
		cursor: pointer;
		transition: background .15s linear;
	}
	.list-item, .req, .res {
		box-shadow: 0px 2px 5px 0px rgba(0, 0, 0, 0.1);
	}
	.list-item.selected {
		background: var(--list-item-sel-bg);
	}

	.GET    { color: var(--method-get) }
	.POST   { color: var(--method-post) }
	.PUT    { color: var(--method-put) }
	.PATCH  { color: var(--method-patch) }
	.DELETE { color: var(--method-delete) }
	.ok     { color: var(--status-ok) }
	.warn   { color: var(--status-warn) }
	.error  { color: var(--status-error) }

	.method { font-size: 0.7em; text-align: left; }
	.status { font-size: 0.8em; }
	.path   { font-size: 0.8em; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; direction: rtl }
	.time   { font-size: 0.7em; color: var(--disabled) }

	pre {
		word-break: normal; word-wrap: break-word; white-space: pre-wrap;
		z-index: 2;
		padding: 1rem;
		width: 100%;
		font-family: inherit;
		font-weight: 400;
		line-height: 1.2em;
	}

	.corner {
		position: absolute;
		top: 0;
		right: 0;
		width: 80px;
		height: 50px;
		background: var(--bg);
		color: var(--disabled);
		display: grid;
		align-content: end;
		justify-content: center;
		transform: rotate(45deg) translate(10px, -40px);
		padding-bottom: 4px;
		font-size: .8em;
		user-select: none;
	}

	.controls {
		display: grid;
		grid-template-columns: repeat(5, 1fr);
		grid-gap: .5rem;
		justify-content: start;
	}
	button {
		background: var(--btn-bg);
		border: 0;
		padding: .5rem 1rem;
		font-size: .75em;
		font-family: inherit;
		color: var(--links);
		cursor: pointer;
		outline: 0;
	}
	button:disabled {
		color: var(--disabled);
		cursor: default;
	}
	button:hover:enabled {
		background: var(--btn-hover);
	}

	.welcome {
		display: grid;
		position: absolute;
		background: rgba(0, 0, 0, .5);
		justify-content: center;
		line-height: 1.5rem;
		z-index: 9;
		color: #fff;
		font-size: 2em;
		top: 50%;
		right: 1rem;
		left: 1rem;
		transform: translate(0%, -50%);
		padding: 3rem;
		box-shadow: 0px 0px 20px 10px rgba(0, 0, 0, 0.1);
		word-break: break-word;
	}
	.welcome span {
		font-size: .5em;
		color: #999;
	}

	@media only screen and (max-width: 1024px) {
		.dashboard { grid-template-columns: .7fr 1fr; grid-template-rows: 1fr 1fr }
		.list { grid-row: 1 / 3 }
		.req { grid-column: 2 }
		.res { grid-column: 2; grid-row: 2 }
		.welcome { font-size: 1.5em }
	}
	@media only screen and (max-width: 484px) {
		.dashboard { grid-template-columns: 1fr; grid-template-rows: 1fr 1fr 1fr; column-gap: 0 }
		.list { grid-area: 1 / 2 }
		.req { grid-row: 2 }
		.res { grid-row: 3 }
	}
	</style>
</head>
<body>

<div class="dashboard" id="app" v-cloak>

	<div class="list">
		<div class="controls">
			<button :disabled="items.length == 0" @click="clearDashboard">clear</button>
		</div>
		<div class="list-inner">
			<div class="list-item" v-for="item in items" :key="item.id" @click="show(item)"
				 :class="{selected: selectedItem.id == item.id}">
				<span class="method" :class="item.method">{{item.method}}</span>
				<span class="path">&lrm;{{item.path}}&lrm;</span>
				<span class="time">{{item.elapsed}}ms</span>
				<span class="status" :class="statusColor(item)">{{item.status == 999 ? 'failed' : item.status}}</span>
			</div>
		</div>
	</div>

	<div class="req">
		<div class="controls">
			<button :disabled="!canPrettifyBody('request')" @click="prettifyBody('request')">prettify</button>
			<button :disabled="selectedItem.id == null" @click="copyCurl">curl</button>
			<button :disabled="selectedItem.id == null" @click="retry">retry</button>
		</div>
		<div class="req-inner">
			<div class="corner">req</div>
			<pre>{{selectedItem.request}}</pre>
		</div>
	</div>

	<div class="res">
		<div class="controls">
			<button :disabled="!canPrettifyBody('response')" @click="prettifyBody('response')">prettify</button>
		</div>
		<div class="res-inner">
			<div class="corner">res</div>
			<pre :class="{error: selectedItem.status == 999}">{{selectedItem.response}}</pre>
		</div>
	</div>

	<div class="welcome" v-show="items.length == 0">
		<p>Waiting for requests on http://localhost:<<.ProxyPort>>/<br>
		<span>Proxying <<.TargetURL>></span></p>
	</div>

</div>

<script type="text/javascript">
var app = new Vue({
	el: '#app',
	data: {
		items: [],
		selectedItem: {},
	},
	created() {
		this.setupStream();
	},
	methods: {
		setupStream() {
			let es = new EventSource(window.location.href + '/conn/');
			es.addEventListener('captures', event => {
				this.items = JSON.parse(event.data).reverse();
			});
			es.onerror = () => {
				this.items = [];
				this.selectedItem = {};
			};
		},
		async show(item) {
			this.selectedItem = { ...this.selectedItem, id: item.id, status: item.status };
			let resp = await fetch(window.location.href + '/info/' + item.id);
			let data = await resp.json();
			this.selectedItem = { ...this.selectedItem,  ...data };
		},
		statusColor(item) {
			if (item.status < 300) return 'ok';
			if (item.status < 400) return 'warn';
			return 'error';
		},
		async clearDashboard() {
			this.selectedItem = {};
			await fetch(window.location.href + '/clear/');
		},
		canPrettifyBody(name) {
			if (!this.selectedItem[name]) return false;
			return this.selectedItem[name].indexOf('Content-Type: application/json') != -1;
		},
		prettifyBody(key) {
			let regex = /\n([\{\[](.*\s*)*[\}\]])/;
			let data = this.selectedItem[key];
			let match = regex.exec(data);
			let body = match[1];
			let prettyBody = JSON.stringify(JSON.parse(body), null, '    ');
			this.selectedItem[key] = data.replace(body, prettyBody);
		},
		copyCurl() {
			let e = document.createElement('textarea');
			e.value = this.selectedItem.curl;
			document.body.appendChild(e);
			e.select();
			document.execCommand('copy');
			document.body.removeChild(e);
		},
		async retry() {
			await fetch(window.location.href + '/retry/' + this.selectedItem.id);
			this.show(this.items[0]);
		},
	},
});
</script>
</body>
</html>`
