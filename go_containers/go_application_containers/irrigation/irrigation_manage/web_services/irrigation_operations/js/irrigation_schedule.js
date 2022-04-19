var working_scheduling;
function start_schedule_select(choice) {
    working_scheduling = choice;
    queue_schedule_start();
}
function queue_schedule_start() {
    hide_all_sections();
    fill_in_schedule_table();
    show_section("queue_schedule");
}
function queue_schedule_init(main_controller, sub_controller, master_flag) {
    // attach select handler 
    create_table("#schedule_table", ["Ref Point", "Select Point", "Step", "Time", "Valves"]);
    attach_button_handler("#schedule_cancel_id", schedule_cancel_idl);
    var select_array = ["select entry", "Queue Entries", "Move Entries"];
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
    $("#schedule_action_select")[0].selectedIndex = 0;
    if (index == 0) {
        return;
    }
    alert(choice);
}
function fill_in_schedule_table() {
    var steps = schedule_step_map[working_scheduling];
    var table_data = [];
    for (i = 0; i < steps.length; i++) {
        var temp = [];
        temp.push(radio_button_element("Action_display_list_select" + i));
        temp.push(check_box_element("Action_display_list_checkbox" + i));
        temp.push(i + 1);
        temp.push(steps[i]["time"]);
        temp.push(steps[i]["steps"]);
        table_data.push(temp);
    }
    load_table("#schedule_table", table_data);
}
