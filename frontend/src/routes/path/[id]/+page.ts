import {GetPath} from "../../../actions/DefaultPathAction.svelte"
import type {Path} from "../../../models"

/** @type {import('./$types').PageLoad} */
export async function load({ params }) {

    let p: Path = await GetPath(params.id)

    return p
  }