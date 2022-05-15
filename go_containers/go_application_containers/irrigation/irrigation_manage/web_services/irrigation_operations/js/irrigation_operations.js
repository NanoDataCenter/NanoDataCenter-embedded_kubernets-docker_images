action_data_list = [];
schedule_data = {};
schedule_name_list = [];
schedule_description_list = [];
schedule_map = [];
function main_form_start() {
    hide_all_sections();
    show_section("main_form");
}
function main_form_init() {
    controller_init();
    attach_button_handler("#manage_select", open_queue_manage);
    attach_button_handler("#manage_valve_group_io", open_valve_group_manage);
    attach_button_handler("#manage_direct_io", station_channel_manage);
}
function load_new_data() {
    var data = {};
    var master_flag = $("#master_controller_select").is(':checked');
    var master_name = $("#master_server").val();
    var sub_name = $("#sub_server").val();
    if (master_flag == true) {
        g_server_key = "true~" + master_name + "~" + sub_name;
    }
    else {
        g_server_key = "false~" + master_name + "~" + sub_name;
    }
    if ($("#master_controller_select").is(':checked') == true) {
        schedule_map = {};
        jquery_populate_select('#irrigation_schedule_select', [], []);
        get_action_data();
        return;
    }
    var data = {};
    data["server_key"] = g_server_key;
    ajax_post_get(ajax_get_schedule, data, ajax_get_schedule_function, "Schedule Data Not Loaded");
}
function ajax_get_schedule_function(data) {
    schedule_data = data;
    console.log("ajax_get_schedule", data);
    // generatate schecu
    schedule_name_list = [];
    schedule_description_list = [];
    schedule_map = {};
    schedule_step_map = {};
    var temp = [];
    for (var i = 0; i < data.length; i++) {
        var name_1 = data[i]["name"];
        temp.push(name_1);
        schedule_map[name_1] = data[i];
    }
    temp.sort();
    console.log("schedule_map", schedule_map);
    schedule_description_list = ["select schedule"];
    schedule_name_list = ["blank"];
    for (var i = 0; i < temp.length; i++) {
        var name_2 = temp[i];
        schedule_name_list.push(name_2);
        var description = schedule_map[name_2]["description"];
        schedule_description_list.push(name_2 + "  :  " + description);
        schedule_step_map[name_2] = process_schedule_step(schedule_map[name_2].steps);
    }
    console.log("schedule step map ", schedule_step_map);
    jquery_populate_select('#irrigation_schedule_select', schedule_name_list, schedule_description_list, show_schedule_page);
    get_action_data();
}
function get_action_data() {
    var data = {};
    data["server_key"] = g_server_key;
    ajax_post_get(ajax_get_actions, data, ajax_process_action_data, "Irrigation Action Data Not Loaded");
}
function ajax_process_action_data(data) {
    action_data_list = [];
    data.sort();
    action_data_list = ["select action"];
    for (var i = 0; i < data.length; i++) {
        action_data_list.push(data[i]);
    }
    jquery_populate_select('#action_select', action_data_list, action_data_list, queue_action_data);
}
function open_queue_manage() {
    alert("open queue manager");
}
function open_valve_group_manage() {
    valve_group_components_start();
}
function station_channel_manage() {
    station_channel_start();
}
function queue_action_data() {
    "action_select";
    index = $("#action_select")[0].selectedIndex;
    choice = $("#action_select").val();
    if (index == 0) {
        return;
    }
    $("#action_select")[0].selectedIndex = 0;
    var data = {};
    data["key"] = g_server_key;
    data["action"] = choice;
    var url_path = "ajax/irrigation/irrigation_manage/post_action";
    ajax_post_confirmation(url_path, data, "Queue Selected Action  " + choice, "Action Queued", "Action Not Queued");
}
function queue_schedule_data(schedule_data) {
    var data = {};
    data["key"] = g_server_key;
    data["schedule"] = schedule_data;
    var url_path = "ajax/irrigation/irrigation_manage/post_schedule";
    ajax_post_confirmation(url_path, data, "Queue Schedule", "Schedule is queue", "Schedule Not Queue");
}
function show_schedule_page() {
    var index = $("#irrigation_schedule_select")[0].selectedIndex;
    var choice = $("#irrigation_schedule_select").val();
    $("#irrigation_schedule_select")[0].selectedIndex = 0;
    if (index == 0) {
        return;
    }
    start_schedule_select(choice);
}
function process_schedule_step(step_data) {
    var return_value = [];
    for (var i = 0; i < step_data.length; i++) {
        var temp = {};
        temp["step"] = i + 1;
        temp["time"] = step_data[i]["time"];
        temp["steps"] = JSON.stringify(process_valve_data(step_data[i]["station"]));
        return_value.push(temp);
    }
    console.log(return_value);
    return return_value;
}
function process_valve_data(station_data) {
    return_value = [];
    stations = Object.keys(station_data);
    for (var i = 0; i < stations.length; i++) {
        var station = stations[i];
        var temp = station_data[station];
        var io_list = Object.keys(temp);
        for (var j = 0; j < io_list.length; j++) {
            return_value.push(station + ":" + io_list[j]);
        }
    }
    return return_value;
}
