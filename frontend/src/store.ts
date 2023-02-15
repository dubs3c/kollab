import { writable } from 'svelte/store';
import type { Path } from './models';


interface Settings {
    DefaultPaths: Path[]
  };

const initialState: Settings = {
    DefaultPaths: []
}

export const defaultPaths = writable(initialState);

/*
export const defaultPaths = () => {
    // creates a new writable store populated with some initial data
    const { subscribe, update, set } = writable(initialState);
  
    return {
        subscribe,
        update,
        set,
        setPath: (url: string) => 
            update(state => (state = {...state, Path: url})),
    };
  };*/