var working_scheduling;
var working_data;
function start_schedule_select(choice) {
    working_scheduling = choice;
    working_data = schedule_step_map[working_scheduling];
    console.log("start schedule", choice, working_scheduling, working_data);
    queue_schedule_start();
}
function queue_schedule_start() {
    hide_all_sections();
    fill_in_schedule_table();
    show_section("queue_schedule");
}
function queue_schedule_init(main_controller, sub_controller, master_flag) {
    // attach select handler 
    Table_create_table("#schedule_table", ["Ref Point", "Select Point", "Index", "Step", "Time", "Valves"], [Table_radio_button_element, Table_check_box_element, Table_to_string, Table_to_string, Table_to_string, Table_to_json]);
    attach_button_handler("#schedule_cancel_id", schedule_cancel_idl);
    var select_array = ["select entry", "Queue All Entries", "Queue Selected Entries", "Change Time", "Move Entries", "Select All", "UnSelect All", "Reset"];
    jquery_populate_select("#schedule_action_select", select_array, select_array, schedule_select_handler);
}
function schedule_cancel_idl() {
    start_section("main_form");
}
function schedule_select_handler() {
    var index;
    var choice;
    index = $("#schedule_action_select")[0].selectedIndex;
    choice = $("#schedule_action_select").val();
    var length = working_data.length;
    if (index == 0) {
        return;
    }
    if ($("#schedule_action_select")[0].selectedIndex == 1) {
        queue_schedule_data(working_data);
    }
    if ($("#schedule_action_select")[0].selectedIndex == 2) {
        var new_indexes = Table_find_check_box_elements("Sched_display_list_checkbox", length);
        var new_table = Table_construct_table_elements(working_data, new_indexes);
        queue_schedule_data(new_table);
    }
    if ($("#schedule_action_select")[0].selectedIndex == 3) {
        var new_indexes = Table_find_check_box_elements("Sched_display_list_checkbox", length);
        if (new_indexes.length == 0) {
            alert("need to select at least one checkbox");
            return;
        }
        step_time_activate_function(time_return_function, time_change_function);
    }
    if ($("#schedule_action_select")[0].selectedIndex == 4) {
        new_indexes = Table_do_move(length, "Sched_display_list_select", "Sched_display_list_checkbox");
        working_data = Table_remap_table(working_data, new_indexes);
        fill_in_schedule_table();
    }
    if ($("#schedule_action_select")[0].selectedIndex == 5) {
        Table_select_check_box_elements("Sched_display_list_checkbox", length);
    }
    if ($("#schedule_action_select")[0].selectedIndex == 6) {
        Table_unselect_check_box_elements("Sched_display_list_checkbox", length);
    }
    if ($("#schedule_action_select")[0].selectedIndex == 7) {
        working_data = schedule_step_map[working_scheduling];
        fill_in_schedule_table();
    }
    $("#schedule_action_select")[0].selectedIndex = 0;
}
function fill_in_schedule_table() {
    var steps = working_data;
    //console.log("steps",steps)
    Table_clear_table("#schedule_table");
    for (var i = 0; i < steps.length; i++) {
        var temp = [];
        temp.push("Sched_display_list_select" + i);
        temp.push("Sched_display_list_checkbox" + i);
        temp.push(i + 1);
        temp.push(steps[i]["step"]);
        temp.push(steps[i]["time"]);
        temp.push(steps[i]["steps"]);
        Table_add_row("#schedule_table", temp);
    }
    Table_load_table("#schedule_table");
}
function time_return_function() {
    queue_schedule_start();
}
function time_change_function(new_time) {
    var new_indexes = Table_find_check_box_elements("Sched_display_list_checkbox", working_data.length);
    for (i = 0; i < new_indexes.length; i++) {
        var index = new_indexes[i];
        working_data[index]["time"] = new_time;
    }
    queue_schedule_start();
}
