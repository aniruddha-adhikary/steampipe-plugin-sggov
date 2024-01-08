# SingStat Plugin for Singapore Government Public Data

Use SQL to query data from [Singapore Department of Statistics](https://singstat.gov.sg) and [Data.gov.sg](https://data.gov.sg).

![Screenshot of SingStat for Steampipe in Action](./dev_docs/cover.png)

> Note: This project is not officially endorsed or sanctioned by the Government of Singapore, Department of
> Statistics or any government agency.

## Quick start

Install the plugin with [Steampipe](https://steampipe.io). Ensure you
have [Steampipe installed and working](https://steampipe.io/downloads)!

1. Clone the code.
    ```shell
    git clone https://github.com/aniruddha-adhikary/steampipe-plugin-sggov.git
    cd steampipe-plugin-sggov
    ```

2. Build and install the plugin using the Makefile.
    ```shell
    make install
    ```

3. Run Steampipe.
   ```shell
   steampipe query
   ```

4. Run a query.
   ```sql
   SELECT id, title
   FROM singstat
   WHERE keyword = 'dialect' AND "searchOption" = 'title';
   ```

## Feature Map

| Feature        | Data.gov.sg | SingStat |
|----------------|-------------|----------|
| Dataset Lookup | ❌           | ✅        |
| Data Fetching  | ❌           | ❌        |

## Developing

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)
