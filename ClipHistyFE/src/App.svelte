<script>
  const fetchImage = (async () => {
    const response = await fetch("http://localhost:3000/hello");
    return await response.json();
  })();
</script>

<style>
  table {
    margin-left: 10%;
    size: 50px;
    width: 60%;
    border-collapse: separate;
    border-spacing: 15px 20px;
  }
  tr > th {
    font-size: 2em;
  }
  tr > td {
    padding-top: 1em;
    padding-bottom: 1em;
    padding-left: 1em;
    border-radius: 3px;
    box-shadow: rgba(15, 15, 15, 0.2) 0px 0px 0px 1px,
      rgba(15, 15, 15, 0.2) 0px 2px 4px;
    background-color: rgb(63, 68, 71);
  }

  th {
    text-align: left;
  }

  div {
    height: 100%;
    width: 100%;
    justify-content: center;
    align-items: center;
    color: rgba(255, 255, 255, 0.9);
    font: BlinkMacSystemFont;
    background-color: #2f3437;
  }

  .date {
    width: 15%;
  }
</style>

<div>
  <table>
    <thead>
      <tr>
        <th class="date">date</th>
        <th>entry</th>
      </tr>
    </thead>
    <tbody>
      {#await fetchImage}
        <p>...waiting</p>
      {:then data}
        {#each data as d}
          <tr>
            <td class="date">{d.Timestamp}</td>
            <td>{d.Content}</td>
          </tr>
        {/each}
      {:catch error}
        <p>An error occurred!</p>
      {/await}

    </tbody>
  </table>
</div>
