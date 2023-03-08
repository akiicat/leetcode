from mkdocs.utils import nest_paths
from natsort import natsorted
import pprint
pp = pprint.PrettyPrinter(indent=4)

def my_sort(data, files, config):
    if isinstance(data, dict):
        for key, value in data.items():
            data[key] = my_sort(value, files, config)
        return data
    elif isinstance(data, list):
        str_arr = []
        obj_arr = []
        for item in data:
            if isinstance(item, str):
                str_arr.append(item)
            else:
                obj_arr.append(my_sort(item, files, config))
        obj_arr = natsorted(obj_arr, key=lambda x: list(x)[0])
        str_arr = natsorted(str_arr)
        data = str_arr + obj_arr
        # data = natsorted(data, key=lambda x: x if isinstance(x, str) else list(x)[0])
        return data
    return data

def on_files(files, config):
    nav_config = config['nav'] or nest_paths(f.src_uri for f in files.documentation_pages())
    # pp.pprint(nav_config)
    config['nav'] = my_sort(nav_config, files, config)
    # pp.pprint(config['nav'])
    return
